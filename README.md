# Telegram bot template
It's been a week since I've started learning Go and I decided to share my way of making a telegram bot using webhook and deploying it on [Yandex Cloud](https://cloud.yandex.com/en-ru/)

## Steps
1. Head to the [BotFather](https://telegram.me/BotFather) and create your bot
2. Now when you got your bot token, paste it in `sendMessage` function in `bot.go`
3. From here you can choose what to use (Heroku, AWS, Yandex Cloud, etc). I will use Yandex Cloud as they give you a [initial grant](https://cloud.yandex.com/en-ru/docs/getting-started/usage-grant)
  - You need to register and make a serverless cloud function (it's the same as AWS Lambda)
  - Put your code in the code editor and make sure your entrypoint is set at `index.Handler` if your file is named index
  - When you checked everything it's time to run it. Press `Create Version` and wait
  - Once status is `Active` **make the function _public_**. That's important for all cloud providers
4. As we are using telegram webhook you need to tell telegram where to send updates:
  ```sh 
  https://api.telegram.org/bot<TOKEN GOES HERE>/setWebHook?url=<FUNCTION LINK GOES HERE>
  ```
  (just paste it in your browser and you will see a reply JSON that the webhook is set)
  
## Usage
In `bot.go` you will find a switch statement that checks if a message contains a `/statement` for the bot. I was making my university schedule bot with several functions and packages but for the template I only left an example function `coolify` so I hope you will figure out how it works ;)

When you're done creating `/statements` you can add them to the bot with the help of the [BotFather](https://telegram.me/BotFather) with `/setcommands` so it will be easier to use the bot.

## Contact
I tried to make it simple and easier to understand but if you have any questions I'll be happy to help! Just email me: danyakurach@gmail.com
  
## Copyright & License
- Copyright (©) 2022 by [Danila Kurach](https://github.com/dupreehkuda)
- Licensed under the terms of the [MIT License](./LICENSE)
