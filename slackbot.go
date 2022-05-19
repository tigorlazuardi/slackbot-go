// slackbot is a library to send messages to slack in simplified form.
//
// It trades customizablity for simplified and scalable api. This library is expected to be used as a tool to send or report messages to channels. Not as an interactive bot.
//
// Scalable means that this library will try to play very nice with slack's api limit and offer back pressure and message drop support to deal with multitude streams of data it expects to receive.
//
// slackbot by default will try to not spam the target channel and uses increasing debounce to not spam the channel with repeated messages.
package slackbot
