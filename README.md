<div align="center">
  <h1>Call Bell</h1>
  <br />
  <a href="https://github.com/ImOverlord/call-bell/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/ImOverlord/call-bell/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  .
  <a href="https://github.com/ImOverlord/call-bell/issues/new?assignees=&labels=question&template=04_SUPPORT_QUESTION.md&title=support%3A+">Ask a Question</a>
</div>

<div align="center">
<br />

[![Project license](https://img.shields.io/github/license/ImOverlord/call-bell.svg?style=flat-square)](LICENSE)

[![Pull Requests welcome](https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=flat-square)](https://github.com/ImOverlord/call-bell/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![code with love by ImOverlord](https://img.shields.io/badge/%3C%2F%3E%20with%20%E2%99%A5%20by-ImOverlord-ff1414.svg?style=flat-square)](https://github.com/ImOverlord)

</div>

---

# About

A simple web application that provides a customizable call bell interface with notification capabilities. It supports multiple notification backends (Slack and Mattermost) and can be easily configured to display custom buttons with different messages. When a button is clicked, it not only sends a notification but also displays a random joke on the interface.

## Prerequisites

- Go 1.21 or later
- Docker (optional, for containerized deployment)
- A notification service (Slack or Mattermost webhook URL)

## Installation

### Docker

The application is available as a Docker image for both AMD64 and ARM64 architectures:

```bash
docker pull ghcr.io/imoverlord/call-bell:latest
```

## Configuration

1. Create a `config.json` file based on the example configuration:
```json
{
    "buttons": [
        {
            "type": "attention",
            "text": "✨Attention✨",
            "message": "Attention has been requested"
        },
        {
            "type": "hungry",
            "text": "😮Hungry😋",
            "message": "Food has been requested"
        }
        // Add more buttons as needed
    ]
}
```

2. Set the required environment variables:
```bash
export NOTIFICATION_URL="your-webhook-url"
export NOTIFICATION_TYPE="slack"  # or "mattermost"
export PORT="8080"  # or your preferred port
```

## Usage

### Running the Application

1. Start the server:

```bash
docker run -d \
  -p 8080:8080 \
  -e NOTIFICATION_URL="your-webhook-url" \
  -e NOTIFICATION_TYPE="slack" \
  -e PORT="8080" \
  -v $(pwd)/config.json:/config.json \
  ghcr.io/imoverlord/call-bell:latest
```

2. Access the web interface at `http://localhost:8080`

### Features

- Support for Slack and Mattermost notifications
- Simple and intuitive web interface
- Built-in random joke responses when buttons are clicked
- Customizable button messages and emojis
- Metrics endpoint for monitoring
- Multi-architecture Docker support

### Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `NOTIFICATION_URL` | URL where notifications will be sent | Yes | - |
| `NOTIFICATION_TYPE` | Type of notification service (mattermost/slack) | Yes | - |
| `PORT` | Port number the application will listen on | No | 9000 |

### API Endpoints

- `GET /`: Main web interface
- `POST /notify`: Send notifications (requires `type` query parameter)
- `GET /metrics`: Prometheus metrics endpoint
- `GET /health`: Health check endpoint (returns 200 OK when service is healthy)

## Development

### Building Docker Images

The project uses GoReleaser for building and publishing Docker images. To build locally:

```bash
export REGISTRY=ghcr.io
export IMAGE_NAME=call-bell
goreleaser release --snapshot --skip=publish --clean
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

