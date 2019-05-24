package wxcrypt

import "testing"

func TestSHA1(t *testing.T) {

	token := "pamtest"
	timestamp := "1409304348"
	nonce := "xxxxxx"
	encrypt := "ZNCXjvcr100slZD4tESc/47itSOLQtHsUMUG9UTfehtAWwGPj4S2yk31dHZmVBwwLuiw60YXpDrnYsrwVnjVyOFozGV5AnCTWGCcFTvNaKnyVW5vWOJGDH2I1mMwafsyHuVrCWcK611hjpLGiVb7aTh6HKugcfmEQTA7IJVMSqdm1L84fM/sduvrjbmGmajIfbuatY7pEyEbPcT6+xW4sNSj16cwIzr6FE+vZXGS2wfnrmt34PKYq64PRCFQ4o3E0su3Pvho/RIR7/rLFeT/Q4emaaL1EP+rmSFrXqZAFSHNkq5EIYGAiFNSGjB/xXsRacLgSyk3BVh3FtzBwGWo7pWm3xdt6wVnxctGkPl8ultTsgsHmFcOGDCx86+3OaI+5FhLlLdy8/7Ho71XmRiqKXU8ZVghiYVpN3b3PdejcMVP7uSltmAoCYosXU6F3lkLEd05m6r6bnJSTj9wkIjuKnvErA/1Sf8WFxEZGtlPWb1ApNDNK1coqocTtiyGfKcov5aZ9p9Tl2tkZHrgZqBIgMcETHdbXvVqd+i93UWz3XRrLBz59ZZKOr99A/ZSbYrAX6i1TWmAcPEallHVwX4gad0/7dmBca0236VTccVXNaqcVnj66nt78gSdSCfyJlNF"

	sha1 := GetSHA1(token, timestamp, nonce, encrypt)
	t.Log(sha1)
}

func BenchmarkSHA1(b *testing.B) {

	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		token := "pamtest"
		timestamp := "1409304348"
		nonce := "xxxxxx"
		encrypt := "ZNCXjvcr100slZD4tESc/47itSOLQtHsUMUG9UTfehtAWwGPj4S2yk31dHZmVBwwLuiw60YXpDrnYsrwVnjVyOFozGV5AnCTWGCcFTvNaKnyVW5vWOJGDH2I1mMwafsyHuVrCWcK611hjpLGiVb7aTh6HKugcfmEQTA7IJVMSqdm1L84fM/sduvrjbmGmajIfbuatY7pEyEbPcT6+xW4sNSj16cwIzr6FE+vZXGS2wfnrmt34PKYq64PRCFQ4o3E0su3Pvho/RIR7/rLFeT/Q4emaaL1EP+rmSFrXqZAFSHNkq5EIYGAiFNSGjB/xXsRacLgSyk3BVh3FtzBwGWo7pWm3xdt6wVnxctGkPl8ultTsgsHmFcOGDCx86+3OaI+5FhLlLdy8/7Ho71XmRiqKXU8ZVghiYVpN3b3PdejcMVP7uSltmAoCYosXU6F3lkLEd05m6r6bnJSTj9wkIjuKnvErA/1Sf8WFxEZGtlPWb1ApNDNK1coqocTtiyGfKcov5aZ9p9Tl2tkZHrgZqBIgMcETHdbXvVqd+i93UWz3XRrLBz59ZZKOr99A/ZSbYrAX6i1TWmAcPEallHVwX4gad0/7dmBca0236VTccVXNaqcVnj66nt78gSdSCfyJlNF"

		GetSHA1(token, timestamp, nonce, encrypt)

	}
	b.StopTimer()
}
