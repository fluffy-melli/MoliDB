### MoliDB

MoliDB는 **REST API**를 활용한 **메모리 기반 데이터베이스**입니다.  
DB 구조와 라우터 설계에 대한 자세한 내용은 [architecture.md](architecture.md)에서 확인할 수 있습니다.

---

### ⬇️ 설치 방법

프로젝트 빌드 및 실행 방법은 아래와 같습니다:

> **주의**: 실행 전에 `.env` 파일에서 `SECRET_KEY`와 `API_KEY` 값을 반드시 수정해 주세요.

#### Docker로 실행하기

```sh
$ docker build -t MoliDB .
$ docker run -d -p 7777:7777 MoliDB
```

---

### 라이선스

`MoliDB`는 **MIT License**를 따릅니다. 코드를 수정하거나 배포할 경우, 라이선스 내용을 준수해 주세요.  

Copyright © All rights reserved.