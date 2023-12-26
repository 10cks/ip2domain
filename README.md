# ip2domain

ip2domain: IP 反查域名，利用go重构了一下，爬虫爬的是 138.com。保存结果为json格式，可集成到 fofaEX中，也可独立运行。

FofaEX：https://github.com/10cks/fofaEX

独立运行命令：

```
.\main.exe -f .\ip.txt -o .\result.json
```

集成进 fofaEX 进行自动化操作：

https://private-user-images.githubusercontent.com/47177550/292841325-1d321f34-d5ce-4ac8-8006-3c93727c9409.mp4?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTEiLCJleHAiOjE3MDM1ODE3MDAsIm5iZiI6MTcwMzU4MTQwMCwicGF0aCI6Ii80NzE3NzU1MC8yOTI4NDEzMjUtMWQzMjFmMzQtZDVjZS00YWM4LTgwMDYtM2M5MzcyN2M5NDA5Lm1wND9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFJV05KWUFYNENTVkVINTNBJTJGMjAyMzEyMjYlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjMxMjI2VDA5MDMyMFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTdiOGVkZDEwOGNlNDI5NzBmYTI2NjUxY2M1MDdkYjk5MDU1NDg4NzQyOTIxNzgwY2RlZTIwZjNiYzQ4Yzc4ZGYmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.d_Wwc7dF-k1rZP3CrYAA3l4w4kiu6nwLkGx24H8yHWA
