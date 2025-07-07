# <img src="https://dashboard.snapcraft.io/site_media/appmedia/2022/03/K6-logo_1.jpg.png" alt="xk6-kafka logo" style="height: 32px; width:32px;"/> xk6-otp

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/BrunoKrugel/xk6-otp/test.yml?branch=main&logo=github)](https://github.com/BrunoKrugel/xk6-otp/actions) [![Coverage Status](https://coveralls.io/repos/github/BrunoKrugel/xk6-otp/badge.svg?branch=main)](https://coveralls.io/github/BrunoKrugel/xk6-otp?branch=main) [![Go Reference](https://pkg.go.dev/badge/github.com/BrunoKrugel/xk6-otp.svg)](https://pkg.go.dev/github.com/BrunoKrugel/xk6-otp)

The `xk6-otp` project is a [k6 extension](https://k6.io/docs/extensions/guides/what-are-k6-extensions/) that enables k6 users to fetch the last OTP code received in Gmail filtered by sender.


Run:

```bash
xk6 build --with github.com/BrunoKrugel/xk6-otp@latest
```

Then run the custom binary with:

```bash
./k6 run build/test.js
```

## Usage

Fetch the last OTP code received in Gmail filtered by sender.

```javascript
// @ts-ignore
import Otp from 'k6/x/otp';

const [message, error] = Otp.lastOtpCode(
    'user@email.com',
    'app password',
    'sender@email.com'
);

if (error) {
    // handle error
}

console.log(message.code)
console.log(message.subject)
console.log(message.date)
console.log(message.sender)
```
