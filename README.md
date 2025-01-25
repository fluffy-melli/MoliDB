프로젝트 구조는 [architecture.md](architecture.md) 에서 확인할 수 있으며 프로젝트 빌드 및 실행 방법은 다음과 같습니다:

>docker
```sh
$ docker build -t MoliDB .
$ docker run -d -p 7777:7777 MoliDB
```
!실행하기전 .env 파일의 SECRET_KEY 와 API_KEY 를 반드시 수정하기 바랍니다