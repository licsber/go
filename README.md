# licsber / go

**以我之名.**

Self Go Library. -- refactor from licsber/licsber(python3).

```bash
go install github.com/licsber/go/cmd/bilibili-merge@latest
# Send FangTang Message, need L_SCT_KEY
go install github.com/licsber/go/cmd/sct@latest
```

## Feature

```text
0.8.20230724 feat merge bilibili videos.
0.3.0 cmd query-phone finish.
0.2.4 auto load home dir .env file.
0.2.1 lFile.Save() will create directory if not exist.
```

## Environment

use `lErr` package will autoload '.env' file in home dir.

```bash
export L_MONGO_URI="MongoDB URI"

export L_S3_ENDPOINT='https://S3 EndPoint'
export L_S3_REGION='us-east-1'
export L_S3_ACCESS='S3 AccessKey'
export L_S3_SECRET='S3 SecretKey'

export L_SCT_KEY="FangTang SendKey"

export L_MEMO_AK='MemoBird AK'
export L_MEMO_DEVICE_ID='MemoBird deviceID'

export L_WXS_APP_ID='WeChat SmallProgram APP_ID'
export L_WXS_APP_SECRET='WeChat SmallProgram APP_SECRET'
export L_WXS_TOKEN='WeChat SmallProgram Message TOKEN'
export L_WXS_AES_KEY='WeChat SmallProgram Message AES_KEY'

export L_PHONE_DAT_DIR='query-phone phone.dat dir'
```

```bash
go env -w GOPROXY=https://proxy.golang.com.cn,direct
go env -w GOPRIVATE=git.licsber.site
git config --global url."ssh://git@git.licsber.site/".insteadOf "https://git.licsber.site/"
```
