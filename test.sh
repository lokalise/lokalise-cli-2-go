#!/usr/bin/env sh

TOKEN=-tc4fab2a40efb0256a57f92f05fd975f6eb6e0866
CMD=./bin/lokalise

echo General actions
# запуск без параметров, либо с ключом --help, -h
# ожидается вывод справки
{CMD}

# вывод версии, только long
{CMD} --version
sleep 2

#- негативный тест: вызвана команда help без глобального ключа --token
# ожидается вывод ошибки "Error: required flag(s) "token" not set" и справки по команде help
#./bin/lokalise help

## группа Team
########################################################################################################################
echo 1. Team actions

# вывод списка команд, доступных для пользователя
# https://lokalise.com/api2docs/curl/#transition-list-all-teams-get
{CMD} team list ${TOKEN}

sleep 2
## группа Team User
########################################################################################################################
echo 2. Team user actions

# https://lokalise.com/api2docs/curl/#transition-list-all-team-users-get
{CMD} ${TOKEN} team-user list --team-id=193277

# https://lokalise.com/api2docs/curl/#transition-retrieve-a-team-user-get
{CMD} ${TOKEN} team-user retrieve --team-id=193277 --user-id=44627

# обновление роли
# https://lokalise.com/api2docs/curl/#transition-update-a-team-user-put
#{CMD} ${TOKEN} team-user update --team-id=193277 --user-id=44627 --role=admin

# удаление роли