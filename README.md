# Slack-Gif-Bot

Slack-Gif-Bot provides links to gif inside of slack channels.

## Description

The Project aims at providing gifs to the person in a slack channel when tagged. This can be useful in casual conversation by giving a random gif on a said prompt.

## Installation

To use this project, follow the steps below:

1. Install Go from the official Go website: [https://golang.org/](https://golang.org/)
2. Clone the repository: `git clone https://github.com/8rxn/slack-gif-bot.git`
3. Change into the project directory: `cd slack-gif-bot`
4. Get Slack API APP_TOKEN & OAUTH_TOKEN from [Slack](https://api.slack.com/) and add them under SLACK_APP_TOKEN & SLACK_BOT_TOKEN in .env file at `/`
    - Install the Slack app to a relevant workspace while getting the tokens.
5. Get Giphy API Key from [Giphy](https://developers.giphy.com/) and add that under GIPHY_API_KEY in .env file at `/`

## Usage

To Build the project, use the following command:

```shell
go build
```

To Run, use the following Command
```shell
go run main.go
```

### Commands

*Hello:*
```bash 
@slack-giphy Hello, I am <Name>
```
<img src="https://github.com/8rxn/slack-gif-bot/assets/75237697/62eb071c-e214-4f2a-97da-382f1bc0e2a6" width={100} />

---

*Gif with text:*
```bash 
@slack-giphy Give me a gif of <prompt>
```
<img src="https://github.com/8rxn/slack-gif-bot/assets/75237697/64c3af2d-9d97-4857-a3c3-7d74e38ec330" width={100} />

---

*Direct prompt:*
```bash 
@slack-giphy <prompt>
```
<img src="https://github.com/8rxn/slack-gif-bot/assets/75237697/03d82da0-ecff-4bba-a830-f698a5cba30e" width={100} />


## Contributing

Contributions are welcome! If you wish to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature`
3. Make your changes and commit them: `git commit -s -m 'Add some feature'`
4. Push the changes to your forked repository: `git push origin feature/your-feature`
5. Open a pull request in the original repository

--
Alternatively, you can also raise issues for any bugs/ feature requests

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

If you have any questions, suggestions, or feedback, please feel free to reach out to the project maintainer at [rajxryn@gmail.com](mailto:rajxryn@gmail.com).

---

Thank you for using this project! We hope it proves useful for your needs.
