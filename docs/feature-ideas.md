### RabbitMQ

Use queue providers so we can use in memory channel based queue and event system 
OR external driver - RabbitMQ

### ~~Connect different types of database~~

~~Update schema definitions and make different database providers - SQLite3, Postgres,
MySQL. *(its even possible single json file instead of db, since data schema is 
not very complicated)*~~ - *Only Postgres.*  

### Db backup

Export database as an json file. Since we use adapters for API, we can use them for this.

### Webhooks

Register webhooks and send request.

### Pages, notes, markdown editor, documentation

### Emails and custom email templates

### Add background image like in trello (?)

Background image, color or gradient

### timer, pomodoro timer working mode

In that mode - timer will stop and start via pomodoro timer logic and send browser notifications about it's current status Also notifications.

### Browser extension.

Not just a timer. Like an easy way to set a new task ( via screenshot and task name on a fly ).
Imagine you need to review new user interface or write a task or note during meeting

### Cli mode

Terminal ui, full cli mode without any web related things - sockets and http routers. Should looks like LazyGit *( just new adapter in ./app/adapters )* 

### Calendar

Calendar for fast task/reminder/due-to assignments. Also import or sync google/outlook calendars

### documentation

add `man-like` documentation page. Big documentation text with all commands. Should gather that information from different packages and embed ( or not ) into binary. That's need's for terminal mode

### Application Delivery

localhost/docker/self-hosted/windows

- *Described in road map*

### OAuth2.0 integration/roles/permissions

- *Described in road map*

### Simple chat

- *Need extend this idea, describe more. Do we really need this? What is that? Slack replacement or just comments under tasks, or conversations like in github ( which are technically still comments )*

### Plugins

Not sure what is that exactly... I like modularity that plugin architecture provide. New features just includes with new plugins that can be developed separately and load at start-time.
- *Described in road map*

### ~~PyQT6 desktop app~~

- *May be just standalone browser window ( electorn like build for frontend )*
