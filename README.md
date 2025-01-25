### MoliDB

> **MoliDB**는 오픈소스로 제공되는 안전한 메모리 데이터베이스로, **REST API를 통해 쉽게** 데이터를 다룰 수 있습니다.
**모든 데이터는 AES 암호화 방식**으로 송수신되어 **민감한 정보도 안전**하게 처리할 수 있습니다.
자세한 아키텍처 및 기능 설명은 [architecture.md](/md/architecture.md)에서 확인하실 수 있습니다.
클라이언트 코드 예제는 [example.md](/md/example.md)에서 확인할 수 있습니다.

### Languages

- [한국어](/README.md)
- [English](/md/README.en.md)

--- 

### ⬇️ 설치 방법

프로젝트 빌드 및 실행 방법은 아래와 같습니다:

> **주의**: 실행 전에 `.env` 파일에서 `SECRET_KEY`와 `API_KEY` 값을 반드시 수정해 주세요.

#### Docker로 실행하기

```sh
$ docker build -t molidb .
$ docker run -d -p 17233:17233 molidb
```

---

### 라이선스

`MoliDB`는 **MIT License**를 따릅니다. 코드를 수정하거나 배포할 경우, 라이선스 내용을 준수해 주세요.  

Copyright © All rights reserved.