import { Client, Intents } from "discord.js";
import { SlashCommandBuilder } from "@discordjs/builders"
import { REST } from "@discordjs/rest"
import { Routes } from "discord-api-types/v9"
import puppeteer from "puppeteer"

const token = "OTY0NjczMzU2Mjg2Nzk1Nzg2.YloEEg.Xfyrx_jbSCCTIfGnnCxHgmu9DKo";
const clientId = "964673356286795786"
const guildId = "964674357056114759"

const commands = [
    new SlashCommandBuilder()
	.setName('echo')
	.setDescription('Replies with your input!')
	.addStringOption(option =>
		option.setName('input')
			.setDescription('The input to echo back')
			.setRequired(true))
].map(command => command.toJSON());

const rest = new REST({ version: '9' }).setToken(token);
const client = new Client({ intents: [Intents.FLAGS.GUILDS, Intents.FLAGS.GUILD_MEMBERS, Intents.FLAGS.GUILD_MESSAGES] });


rest.put(Routes.applicationGuildCommands(clientId, guildId), { body: commands })
    .then(() => console.log('Successfully registered application commands.'))
    .catch(console.error);



client.once('ready', () => {
    console.log('Ready!');
});

client.on('interactionCreate', async interaction => {
    if (!interaction.isCommand()) return;
    const { commandName } = interaction;
    if (commandName === 'echo') {
        console.log(interaction.options.data)
        await interaction.reply('Pong!');
        (async () => {
            await interaction.followUp('start chrome');
            const browser = await puppeteer.launch();
            const page = await browser.newPage();
            await page.goto('https://www.npmjs.com/package/@types/puppeteer');
            await page.screenshot({ path: 'example.png' });
            await browser.close();
        })();
        // await interaction.followUp('Pong!');
    }
});


client.on("messageCreate", (msg) => {
    //console.log(msg.embeds[0].fields)
})


client.login(token);
