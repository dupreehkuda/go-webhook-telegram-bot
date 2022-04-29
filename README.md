# Telegram bot template
It's been a week since I've started learning Go and I decided to share my way of making a telegram bot using webhook and deploying it on [Yandex Cloud](https://cloud.yandex.com/en-ru/)

[![Deploy](https://yastatic.net/s3/cloud/www/static/freeze/assets/img/logo.54a174a9.svg)](https://cloud.yandex.com/en-ru/services/functions)

## Steps
1. Head to the [BotFather](https://telegram.me/BotFather) and create your bot
2. Now when you got your token, paste it in `sendMessage` function
3. From here you can choose what to use (Heroku, AWS, Yandex Cloud, etc). I will use Yandex Cloud as they give you a [initial grant](https://cloud.yandex.com/en-ru/docs/getting-started/usage-grant)
  - You need to register and make a serverless cloud function (the same as AWS Lambda)
  - Put your code in the code editor and make sure your entrypoint is at `index.Handler` if your file is named index
  - When you checked everything it's time to run it. Press `Create Version` and wait
  - Once status is `Active` **make the function _public_** 
4. As we are using telegram webhook you need to tell telegram where to send updates:
  ```sh 
  https://api.telegram.org/bot<TOKEN GOES HERE>/setWebHook?url=<FUNCTION LINK GOES HERE>
  ```
  
## Copyright & License
- Copyright (©) 2022 by [Danila Kurach](https://github.com/dupreehkuda)
- Licensed under the terms of the [MIT License](./LICENSE)
