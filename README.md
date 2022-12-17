# licsber / go

**以我之名.**

Self Go Library.  -- refactor from licsber/licsber(python3).

```bash
# Send FangTang Message, need L_SCT_KEY
go install github.com/licsber/go/cmd/sct@latest
```

## Environment

```bash
export L_MONGO_URI="MongoDB URI"

export L_S3_ENDPOINT='http://S3 EndPoint'
export L_S3_REGION='us-east-1'
export L_S3_ACCESS='S3 AccessKey'
export L_S3_SECRET='S3 SecretKey'

export L_SCT_KEY="FangTang SendKey"

export L_MEMO_AK='MemoBird AK'
export L_MEMO_DEVICE_ID='MemoBird deviceID'
```

```bash
go env -w GOPROXY=https://proxy.golang.com.cn,direct
```
