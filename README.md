# slslog

## Overview
AWS Lambda の標準出力に吐かれたログを Cloudwatch Logs Insight で trace 付きで一連の処理を追うことのできるための Logger です。

This is inspired by https://github.com/DeNA/aelog.

## Usage

```go
const label = "$LogLabelName"

func main() {
    slslog.SetLogLabel(label)
    span := slslog.StartSpan(context.Background(), label)
    defer span.End()

    ctx := span.Context()
    Infof(ctx, "this is slslog output")

    // Output:
    // {"severity":"INFO","message":"this is slslog output","trace":"service/$LogLabelName/trace/...","span":"service/$LogLabelName/span/..."}
}
```

## See Also

- [サポートされるログと検出されるフィールド](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html)
