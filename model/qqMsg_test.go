package model

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    *QQMsg
		wantErr bool
	}{
		{"basic", args{msg: "{\"Content\":\"大哥哥\",\"GroupPic\":[{\"FileId\":2417873344,\"FileMd5\":\"B+cQNlJ22Y18Y0o+4F03ag==\",\"FileSize\":5917,\"ForwordBuf\":\"Eip7MDdFNzEwMzYtNTI3Ni1EOThELTdDNjMtNEEzRUUwNUQzNzZBfS5qcGciACoEAwAAADJfFTYgOTVrRDFBOTAxZGQxYzBjZGM0YzBiNyAgICAgIDUwICAgICAgICAgICAgICAgIHswN0U3MTAzNi01Mjc2LUQ5OEQtN0M2My00QTNFRTA1RDM3NkF9LmpwZ0FCQ0E4wKP3gAlAt4GT7gxIUFBBWgBgAWoQB+cQNlJ22Y18Y0o+4F03anJaL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvMTk4P3Rlcm09MjU1ggFYL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvMD90ZXJtPTI1NbABjAG4AYwByAGdLtgBjAHgAYwB+gFaL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvNDAwP3Rlcm09MjU1gAKMAYgCjAE=\",\"ForwordField\":8,\"Url\":\"http://gchat.qpic.cn/gchatpic_new/570966274/992486550-2534335053-07E710365276D98D7C634A3EE05D376A/0?vuin=2834323101\\u0026term=255\\u0026pictype=0\"},{\"FileId\":2417873344,\"FileMd5\":\"B+cQNlJ22Y18Y0o+4F03ag==\",\"FileSize\":5917,\"ForwordBuf\":\"Eip7MDdFNzEwMzYtNTI3Ni1EOThELTdDNjMtNEEzRUUwNUQzNzZBfS5qcGciACoEAwAAADILFTcgMTFBREFCQ0E4wKP3gAlAt4GT7gxIUFBBWgBgAWoQB+cQNlJ22Y18Y0o+4F03anJaL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvMTk4P3Rlcm09MjU1ggFYL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvMD90ZXJtPTI1NbABjAG4AYwByAGdLtgBjAHgAYwB+gFaL2djaGF0cGljX25ldy81NzA5NjYyNzQvOTkyNDg2NTUwLTI0MTc4NzMzNDQtMDdFNzEwMzY1Mjc2RDk4RDdDNjM0QTNFRTA1RDM3NkEvNDAwP3Rlcm09MjU1gAKMAYgCjAE=\",\"ForwordField\":8,\"Url\":\"http://gchat.qpic.cn/gchatpic_new/570966274/992486550-2534335053-07E710365276D98D7C634A3EE05D376A/0?vuin=2834323101\\u0026term=255\\u0026pictype=0\"}],\"Tips\":\"[群图片]\"}"},
			&QQMsg{
				Content: "大哥哥",
				UserID:  nil,
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQQMsg(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQQMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQQMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
