# xk6-otp

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/BrunoKrugel/xk6-otp/test.yml?branch=main&logo=github)](https://github.com/BrunoKrugel/xk6-otp/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/BrunoKrugel/xk6-otp.svg)](https://pkg.go.dev/github.com/BrunoKrugel/xk6-otp)

A [k6 extension](https://k6.io/docs/extensions/guides/what-are-k6-extensions/) that fetches the last OTP (one-time password) code from Gmail via IMAP, filtered by sender and optional subject keyword.

## Prerequisites

- [xk6](https://github.com/grafana/xk6) installed (`go install go.k6.io/xk6@latest`)
- A Gmail account with an [App Password](https://support.google.com/accounts/answer/185833) (regular password won't work with IMAP)

## Install

```bash
xk6 build --with github.com/BrunoKrugel/xk6-otp@latest
```

## API

### `Otp.lastOtpCode(email, password, senderEmail, includeFilter)`

Fetches the most recent email from `senderEmail` whose subject contains `includeFilter` (case-insensitive), and extracts the 6-digit OTP code.

**Returns**: `[message, error]`
- `message` — an object with `.code`, `.subject`, `.date`, `.sender`, or `null` on error
- `error` — an error string, or empty on success

### `Otp.lastOtpCodeBySender(email, password, senderEmail)`

Same as above, but without subject filtering — returns the most recent email from the sender regardless of subject.

**Returns**: `[message, error]` (same shape)

## Usage

```javascript
// @ts-ignore
import Otp from 'k6/x/otp';

export default function () {
    const [message, error] = Otp.lastOtpCode(
        'user@email.com',
        'app-password',
        'sender@email.com',
        'verification'
    );

    if (error) {
        console.error('Failed to fetch OTP:', error);
        return;
    }

    console.log('OTP code:', message.code);
    console.log('Subject:', message.subject);
    console.log('Date:', message.date);
    console.log('Sender:', message.sender);
}
```

Then run:

```bash
./k6 run script.js
```

## Notes

- Only Gmail IMAP (`imap.gmail.com:993`) is supported
- OTP extraction matches exactly 6 consecutive digits in the email subject
- The `includeFilter` parameter is compared case-insensitively against the subject line
