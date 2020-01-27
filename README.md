Описание

Для решения задачи взяты 3 приложения:
1. client - для удобства посылки POST строки
2. server - http сервер
3. reserser - переврачивает строку


Компиляция

git clone 
cd server; make; cd -
cd reverser; make; cd -
cd client; make;


Запуск

из дирeктории server запустить 
./server

из директории client запустить
./client -srt=abcd
или
./client -srt=error



