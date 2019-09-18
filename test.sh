#!/usr/bin/env sh

TOKEN=-tc4fab2a40efb0256a57f92f05fd975f6eb6e0866

# запуск без параметров, либо с ключом --help, -h
# ожидается вывод справки
./bin/lokalise
./bin/lokalise --help
./bin/lokalise -h

# вывод версии, только long
./bin/lokalise --version

#- негативный тест: вызвана команда help без глобального ключа --token
# ожидается вывод ошибки "Error: required flag(s) "token" not set" и справки по команде help
./bin/lokalise help

## группа Team
########################################################################################################################

# вывод только справки по группе team
./bin/lokalise team
#- негативный тест: вызвана команда list без глобального ключа --token
./bin/lokalise team list
#- негативный тест: неправильный токен
# ожидается вывод ошибки "Error: API request error 400 Invalid `X-Api-Token` header" и справки по команде team list
./bin/lokalise team list -tasdasd
# вывод списка команд, доступных для пользователя
# ожидается вывод в соотв с документацией
# https://lokalise.com/api2docs/curl/#transition-list-all-teams-get
./bin/lokalise team list ${TOKEN}

## группа Team User
########################################################################################################################

#- негативный тест: вызвана команда list без ключа --team-id
# ожидается вывод ошибки "Error: required flag(s) "team-id" not set" и справки по команде team-user list
./bin/lokalise ${TOKEN} team-user list
# ожидается вывод в соотв с документацией
# https://lokalise.com/api2docs/curl/#transition-list-all-team-users-get
./bin/lokalise ${TOKEN} team-user list --team-id=193277
# ожидается вывод в соотв с документацией
# https://lokalise.com/api2docs/curl/#transition-retrieve-a-team-user-get
./bin/lokalise ${TOKEN} team-user retrieve --team-id=193277 --user-id=44627
# обновление роли
# ожидается вывод в соотв с документацией
# https://lokalise.com/api2docs/curl/#transition-update-a-team-user-put
./bin/lokalise -tc4fab2a40efb0256a57f92f05fd975f6eb6e0866 team-user update --team-id=193277 --user-id=44627 --role=admin