// Copyright (c) 2021 by library authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	tcg "github.com/bluecmd/go-tcg-storage"
	"github.com/bluecmd/go-tcg-storage/drive"
	"github.com/davecgh/go-spew/spew"
)

func TestComID(d tcg.DriveIntf) tcg.ComID {
	comID, err := tcg.GetComID(d)
	if err != nil {
		log.Printf("Unable to auto-allocate ComID: %v", err)
		return tcg.ComIDInvalid
	}
	log.Printf("Allocated ComID 0x%08x", comID)
	valid, err := tcg.IsComIDValid(d, comID)
	if err != nil {
		log.Printf("Unable to validate allocated ComID: %v", err)
		return tcg.ComIDInvalid
	}
	if !valid {
		log.Printf("Allocated ComID not valid")
		return tcg.ComIDInvalid
	}
	log.Printf("ComID validated successfully")

	if err := tcg.StackReset(d, comID); err != nil {
		log.Printf("Unable to reset the synchronous protocol stack: %v", err)
		return tcg.ComIDInvalid
	}
	log.Printf("Synchronous protocol stack reset successfully")
	return comID
}

func TestSession(d tcg.DriveIntf, d0 *tcg.Level0Discovery, comID tcg.ComID) *tcg.Session {
	if comID == tcg.ComIDInvalid {
		log.Printf("Auto-allocation ComID test failed earlier, selecting first available base ComID")
		if d0.OpalV2 != nil {
			log.Printf("Selecting OpalV2 ComID")
			comID = tcg.ComID(d0.OpalV2.BaseComID)
		} else if d0.PyriteV1 != nil {
			log.Printf("Selecting PyriteV1 ComID")
			comID = tcg.ComID(d0.PyriteV1.BaseComID)
		} else if d0.PyriteV2 != nil {
			log.Printf("Selecting PyriteV2 ComID")
			comID = tcg.ComID(d0.PyriteV1.BaseComID)
		} else {
			log.Printf("No supported feature found, giving up without a ComID ...")
			return nil
		}
	}
	log.Printf("Creating control session with ComID 0x%08x\n", comID)
	cs, err := tcg.NewControlSession(d, d0.TPer, tcg.WithComID(tcg.ComID(d0.OpalV2.BaseComID)))
	if err != nil {
		log.Printf("s.NewControlSession failed: %v", err)
		return nil
	}
	spew.Dump(cs)
	// TODO: Move this to a test case instead
	if err := cs.Close(); err != nil {
		log.Fatalf("Test of ControlSession Close failed: %v", err)
	}
	s, err := cs.NewSession()
	if err != nil {
		log.Printf("s.NewSession failed: %v", err)
		return nil
	}
	return s
}

func main() {
	spew.Config.Indent = "  "

	d, err := drive.Open(os.Args[1])
	if err != nil {
		log.Fatalf("drive.Open: %v", err)
	}
	defer d.Close()

	fmt.Printf("===> DRIVE SECURITY INFORMATION\n")
	spl, err := drive.SecurityProtocols(d)
	if err != nil {
		log.Fatalf("drive.SecurityProtocols: %v", err)
	}
	log.Printf("SecurityProtocols: %+v", spl)
	crt, err := drive.Certificate(d)
	if err != nil {
		log.Fatalf("drive.Certificate: %v", err)
	}
	log.Printf("Drive certificate:")
	spew.Dump(crt)
	fmt.Printf("\n")

	fmt.Printf("===> TCG AUTO ComID SELF-TEST\n")
	comID := TestComID(d)
	fmt.Printf("\n")

	fmt.Printf("===> TCG FEATURE DISCOVERY\n")
	d0, err := tcg.Discovery0(d)
	if err != nil {
		log.Fatalf("tcg.Discovery0: %v", err)
	}
	spew.Dump(d0)
	fmt.Printf("\n")

	fmt.Printf("===> TCG SESSION\n")

	s := TestSession(d, d0, comID)
	if s == nil {
		log.Printf("No session, unable to continue")
		return
	}
	spew.Dump(s)

	if err := s.Close(); err != nil {
		log.Fatalf("Session.Close failed: %v", err)
	}
}
