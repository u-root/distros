// Copyright 2016-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var (
	TimeStampNames = map[int]string{
		0:                    "1st timestamp",
		TS_START_ROMSTAGE:    "start of rom stage",
		TS_BEFORE_INITRAM:    "before ram initialization",
		TS_AFTER_INITRAM:     "after ram initialization",
		TS_END_ROMSTAGE:      "end of romstage",
		TS_START_VBOOT:       "start of verified boot",
		TS_END_VBOOT:         "end of verified boot",
		TS_START_COPYRAM:     "starting to load ramstage",
		TS_END_COPYRAM:       "finished loading ramstage",
		TS_START_RAMSTAGE:    "start of ramstage",
		TS_START_BOOTBLOCK:   "start of bootblock",
		TS_END_BOOTBLOCK:     "end of bootblock",
		TS_START_COPYROM:     "starting to load romstage",
		TS_END_COPYROM:       "finished loading romstage",
		TS_START_ULZMA:       "starting LZMA decompress (ignore for x86)",
		TS_END_ULZMA:         "finished LZMA decompress (ignore for x86)",
		TS_DEVICE_ENUMERATE:  "device enumeration",
		TS_DEVICE_CONFIGURE:  "device configuration",
		TS_DEVICE_ENABLE:     "device enable",
		TS_DEVICE_INITIALIZE: "device initialization",
		TS_DEVICE_DONE:       "device setup done",
		TS_CBMEM_POST:        "cbmem post",
		TS_WRITE_TABLES:      "write tables",
		TS_LOAD_PAYLOAD:      "load payload",
		TS_ACPI_WAKE_JUMP:    "ACPI wake jump",
		TS_SELFBOOT_JUMP:     "selfboot jump",

		TS_START_COPYVER:     "starting to load verstage",
		TS_END_COPYVER:       "finished loading verstage",
		TS_START_TPMINIT:     "starting to initialize TPM",
		TS_END_TPMINIT:       "finished TPM initialization",
		TS_START_VERIFY_SLOT: "starting to verify keyblock/preamble (RSA)",
		TS_END_VERIFY_SLOT:   "finished verifying keyblock/preamble (RSA)",
		TS_START_HASH_BODY:   "starting to verify body (load+SHA2+RSA) ",
		TS_DONE_LOADING:      "finished loading body (ignore for x86)",
		TS_DONE_HASHING:      "finished calculating body hash (SHA2)",
		TS_END_HASH_BODY:     "finished verifying body signature (RSA)",

		TS_DC_START:                     "depthcharge start",
		TS_RO_PARAMS_INIT:               "RO parameter init",
		TS_RO_VB_INIT:                   "RO vboot init",
		TS_RO_VB_SELECT_FIRMWARE:        "RO vboot select firmware",
		TS_RO_VB_SELECT_AND_LOAD_KERNEL: "RO vboot select&load kernel",
		TS_RW_VB_SELECT_AND_LOAD_KERNEL: "RW vboot select&load kernel",
		TS_VB_SELECT_AND_LOAD_KERNEL:    "vboot select&load kernel",
		TS_VB_EC_VBOOT_DONE:             "finished EC verification",
		TS_CROSSYSTEM_DATA:              "crossystem data",
		TS_START_KERNEL:                 "start kernel",

		// FSP related timestamps
		TS_FSP_MEMORY_INIT_START:   "calling FspMemoryInit",
		TS_FSP_MEMORY_INIT_END:     "returning from FspMemoryInit",
		TS_FSP_TEMP_RAM_EXIT_START: "calling FspTempRamExit",
		TS_FSP_TEMP_RAM_EXIT_END:   "returning from FspTempRamExit",
		TS_FSP_SILICON_INIT_START:  "calling FspSiliconInit",
		TS_FSP_SILICON_INIT_END:    "returning from FspSiliconInit",
		TS_FSP_BEFORE_ENUMERATE:    "calling FspNotify(AfterPciEnumeration)",
		TS_FSP_AFTER_ENUMERATE:     "returning from FspNotify(AfterPciEnumeration)",
		TS_FSP_BEFORE_FINALIZE:     "calling FspNotify(ReadyToBoot)",
		TS_FSP_AFTER_FINALIZE:      "returning from FspNotify(ReadyToBoot)",
	}

	memTags = map[uint32]string{
		LB_MEM_RAM:         "LB_MEM_RAM",
		LB_MEM_RESERVED:    "LB_MEM_RESERVED",
		LB_MEM_ACPI:        "LB_MEM_ACPI",
		LB_MEM_NVS:         "LB_MEM_NVS",
		LB_MEM_UNUSABLE:    "LB_MEM_UNUSABLE",
		LB_MEM_VENDOR_RSVD: "LB_MEM_VENDOR_RSVD",
		LB_MEM_TABLE:       "LB_MEM_TABLE",
	}
	serialNames = map[uint32]string{
		LB_SERIAL_TYPE_IO_MAPPED:     "IO_MAPPED",
		LB_SERIAL_TYPE_MEMORY_MAPPED: "MEMORY_MAPPED",
	}
	consoleNames = map[uint32]string{
		LB_TAG_CONSOLE_SERIAL8250:    "SERIAL8250",
		LB_TAG_CONSOLE_VGA:           "VGA",
		LB_TAG_CONSOLE_BTEXT:         "BTEXT",
		LB_TAG_CONSOLE_LOGBUF:        "LOGBUF",
		LB_TAG_CONSOLE_SROM:          "SROM",
		LB_TAG_CONSOLE_EHCI:          "EHCI",
		LB_TAG_CONSOLE_SERIAL8250MEM: "SERIAL8250MEM",
	}
	tagNames = map[uint32]string{
		LB_TAG_UNUSED:                "LB_TAG_UNUSED",
		LB_TAG_MEMORY:                "LB_TAG_MEMORY",
		LB_TAG_HWRPB:                 "LB_TAG_HWRPB",
		LB_TAG_MAINBOARD:             "LB_TAG_MAINBOARD",
		LB_TAG_VERSION:               "LB_TAG_VERSION",
		LB_TAG_EXTRA_VERSION:         "LB_TAG_EXTRA_VERSION",
		LB_TAG_BUILD:                 "LB_TAG_BUILD",
		LB_TAG_COMPILE_TIME:          "LB_TAG_COMPILE_TIME",
		LB_TAG_COMPILE_BY:            "LB_TAG_COMPILE_BY",
		LB_TAG_COMPILE_HOST:          "LB_TAG_COMPILE_HOST",
		LB_TAG_COMPILE_DOMAIN:        "LB_TAG_COMPILE_DOMAIN",
		LB_TAG_COMPILER:              "LB_TAG_COMPILER",
		LB_TAG_LINKER:                "LB_TAG_LINKER",
		LB_TAG_ASSEMBLER:             "LB_TAG_ASSEMBLER",
		LB_TAG_VERSION_TIMESTAMP:     "LB_TAG_VERSION_TIMESTAMP",
		LB_TAG_SERIAL:                "LB_TAG_SERIAL",
		LB_TAG_CONSOLE:               "LB_TAG_CONSOLE",
		LB_TAG_FORWARD:               "LB_TAG_FORWARD",
		LB_TAG_FRAMEBUFFER:           "LB_TAG_FRAMEBUFFER",
		LB_TAG_GPIO:                  "LB_TAG_GPIO",
		LB_TAG_VDAT:                  "LB_TAG_VDAT",
		LB_TAG_VBNV:                  "LB_TAG_VBNV",
		LB_TAB_VBOOT_HANDOFF:         "LB_TAB_VBOOT_HANDOFF",
		LB_TAB_DMA:                   "LB_TAB_DMA",
		LB_TAG_RAM_OOPS:              "LB_TAG_RAM_OOPS",
		LB_TAG_MTC:                   "LB_TAG_MTC",
		LB_TAG_TIMESTAMPS:            "LB_TAG_TIMESTAMPS",
		LB_TAG_CBMEM_CONSOLE:         "LB_TAG_CBMEM_CONSOLE",
		LB_TAG_MRC_CACHE:             "LB_TAG_MRC_CACHE",
		LB_TAG_ACPI_GNVS:             "LB_TAG_ACPI_GNVS",
		LB_TAG_WIFI_CALIBRATION:      "LB_TAG_WIFI_CALIBRATION",
		LB_TAG_X86_ROM_MTRR:          "LB_TAG_X86_ROM_MTRR",
		LB_TAG_BOARD_ID:              "LB_TAG_BOARD_ID",
		LB_TAG_MAC_ADDRS:             "LB_TAG_MAC_ADDRS",
		LB_TAG_RAM_CODE:              "LB_TAG_RAM_CODE",
		LB_TAG_SPI_FLASH:             "LB_TAG_SPI_FLASH",
		LB_TAG_BOOT_MEDIA_PARAMS:     "LB_TAG_BOOT_MEDIA_PARAMS",
		LB_TAG_CBMEM_ENTRY:           "LB_TAG_CBMEM_ENTRY",
		LB_TAG_SERIALNO:              "LB_TAG_SERIALNO",
		LB_TAG_CMOS_OPTION_TABLE:     "LB_TAG_CMOS_OPTION_TABLE",
		LB_TAG_OPTION:                "LB_TAG_OPTION",
		LB_TAG_OPTION_ENUM:           "LB_TAG_OPTION_ENUM",
		LB_TAG_OPTION_DEFAULTS:       "LB_TAG_OPTION_DEFAULTS",
		LB_TAG_OPTION_CHECKSUM:       "LB_TAG_OPTION_CHECKSUM",
		LB_TAG_PLATFORM_BLOB_VERSION: "LB_TAG_PLATFORM_BLOB_VERSION",
	}
	tsNames = map[uint32]string{
		TS_DC_START:                     "TS_DC_START",
		TS_RO_PARAMS_INIT:               "TS_RO_PARAMS_INIT",
		TS_RO_VB_INIT:                   "TS_RO_VB_INIT",
		TS_RO_VB_SELECT_FIRMWARE:        "TS_RO_VB_SELECT_FIRMWARE",
		TS_RO_VB_SELECT_AND_LOAD_KERNEL: "TS_RO_VB_SELECT_AND_LOAD_KERNEL",
		TS_RW_VB_SELECT_AND_LOAD_KERNEL: "TS_RW_VB_SELECT_AND_LOAD_KERNEL",
		TS_VB_SELECT_AND_LOAD_KERNEL:    "TS_VB_SELECT_AND_LOAD_KERNEL",
		TS_VB_EC_VBOOT_DONE:             "TS_VB_EC_VBOOT_DONE",
		TS_CROSSYSTEM_DATA:              "TS_CROSSYSTEM_DATA",
		TS_START_KERNEL:                 "TS_START_KERNEL",
	}
)
