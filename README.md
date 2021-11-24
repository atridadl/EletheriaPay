# EleutheriaPay
![EleutheriaPay](https://cloud.atridad.dev/s/jAeYW6nYd7W6sma/preview "EleutheriaPay")

<a href="https://support.atrilahiji.dev/" target="_blank">
    <img src="https://support.atrilahiji.dev/api/badge" />
</a>

Elutheria Pay is a customizable self-hosted FOSS donation page for FOSS developers

DEMO: https://eleutheriapaydemo.atrilahiji.dev

# Requirements
- Node.js 12 or higher (Build)
- Golang 1.16.x or higher (Run)
- A linux machine or WSL if you would like to use the helper scripts I wrote (Build + Run)
- A stripe account with email notifications for transactions enabled: https://stripe.com/docs/receipts (Run)

# Install
## Download the latest release and extract it (sort by name)
https://s3.atridad.dev/minio/eleutheriapay/bin/

## Edit the default .env file that is created (the parameters are outlined below)
IMPORTANT: Please ensure your stripe publishable and secret keys are set before running EleutheriaPay.

## Run the binary
EleutheriaPay will run a server on localhost:3000. Please use this port when configuring a reverse proxy such as NGINX.
```
./eleutheriapay-${VERSION}-${GOOS}-${GOARCH}
```

# Build from source
## Install the rice tool to bundle the frontend with the backend
```
go get github.com/GeertJohan/go.rice/rice
```
## Run the build script
```
./build.sh
```

# Configuraton
## Site Configuration
```
OWNER_NAME: The name of the individual or organization this page is accepting donations for.
META_TITLE: The metadata title for the site.
TITLE: The title on the homepage.
DESCRIPTION: The description of the homepage.
DEFAULT_CURRENCY: The default currency when the page is first reached.

STRIPE_PK: Stripe publishable key.
STRIPE_SK: Stripe secret key.
STRIPE_ACC: Stripe account ID
STRIPE_API_VERSION: Stripe API version
STRIPE_LOCALE: Stripe locale (auto is recommended)

OWNCAST_HOSTNAME: The hostname of your Owncast instance (optional)
OWNCAST_TOKEN: Your Owncast API Token

BADGE_TEXT: The text show on the donation badge.
BADGE_COLOR: The hex value of the donation badge color.
BADGE_FONT_SIZE: An integer value representing the font size of the donationbadge.
BADGE_HEIGHT: The height of the donation badge in pixels.

BACK_LINK: The link the user a redirected to when clicking on the back button.

Social links:
EMAIL_LINK
GIT_LINK
GITHUB_LINK
GITLAB_LINK
FACEBOOK_LINK
TWITTER_LINK
INSTAGRAM_LINK
LINKEDIN_LINK
```

## Theme Configuration
The theme can be customized via the assets directory. Here you can change the logo, favicon, and modify the CSS.

# Embeddable badge
A donate badge can be embedded onto any page or repository to route users to your EleutheriaPay instance. You will need to embed the following endpoint and link it to your instance:

```
<a href="https://donate.<DOMAIN>" target="_blank">
    <img src="https://donate.<DOMAIN>/api/badge" />
</a>
```

# Official Integrations
<a href="https://owncast.online/" target="_blank" rel="nofollow">
    <img src="https://owncast.online/images/logo.svg" width="50" height="50" />
</a>
