# <img src="https://dashboard.snapcraft.io/site_media/appmedia/2022/03/K6-logo_1.jpg.png" alt="xk6-kafka logo" style="height: 32px; width:32px;"/> xk6-otp

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

```javascript
// @ts-ignore
import Otp from 'k6/x/otp';

const [message, error] = Otp.lastOtpCodeBySender(
    'user@email.com',
    'app password',
    'sender@email.com'
);

console.log(message.code)
```
