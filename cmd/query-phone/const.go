package main

//goland:noinspection ALL
const (
	CMCC               byte = iota + 0x01 //中国移动
	CUCC                                  //中国联通
	CTCC                                  //中国电信
	CTCCV                                 //电信虚拟运营商
	CUCCV                                 //联通虚拟运营商
	CMCCV                                 //移动虚拟运营商
	INT_LEN            = 4
	CHAR_LEN           = 1
	HEAD_LENGTH        = 8
	PHONE_INDEX_LENGTH = 9
	PHONE_DAT          = "phone.dat"
)
