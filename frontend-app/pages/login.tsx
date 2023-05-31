import axios from "axios";
import { GetServerSideProps, NextPage } from "next";
import router from "next/router";
import { useState } from "react";
import { serialize } from 'cookie';


type User = {
  LoggedIn: boolean;
  UserId: string;
};
type HomeProps = {
  user: User;
};

const LoginPage: NextPage = () => {
    //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
    // process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (event: React.MouseEvent<HTMLElement>) => {

    axios.post("http://localhost:8080/login", { userId, password }, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      withCredentials: true,
    })
      .then(() => {
        // ログイン成功時の処理
        window.location.href = "/mypage"; // リダイレクト先のURLを指定
      })
      .catch((error) => {
        // ログイン失敗時の処理
        console.error(error);
      });
  };

  return (
    <>
      <h1>LOGIN</h1>
      <form>
        <p>
          ユーザーID：
          <input
            type="text"
            name="userId"
            value={userId}
            onChange={(event) => setUserId(event.target.value)}
            required
            minLength={4}
            maxLength={16}
          />
        </p>
        <p>
          パスワード：
          <input
            type="password"
            name="password"
            value={password}
            onChange={(event) => setPassword(event.target.value)}
            required
            minLength={4}
            maxLength={20}
          />
        </p>
        <p>
          <button type="button" onClick={handleSubmit}>
            ログイン
            </button>
          <button type="button" onClick={() => router.push("/")}>
            戻る
          </button>
        </p>
      </form>
    </>
  );
};

export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
  //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  // process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  // const response = await axios.get("http://localhost:8080/login");
  const response = await axios.get("http://localhost:8080/login", {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      withCredentials: true,
    });
  const user = response.data;

  // // サーバーサイドでクッキーの値を取得
  // const cookieValue = context.req.headers.cookie ?? '';
  // // const decodedCookieValue = Buffer.from(cookieValue, 'base64').toString('utf-8');
  
  // // const cookieParts = cookieValue.split(';'); // クッキーをセミコロンで分割
  // // const formattedCookieValue = cookieParts[1]; // 必要な情報を選択
  

  //   // クッキーを設定
  //   const serializedCookie = serialize('loginUserIdKey', cookieValue, {
  //     // クッキーのオプションを指定
  //     path: '/',
  //     maxAge: 0, // クッキーの有効期限（秒）
  //     httpOnly: false, // クッキーがHTTPプロトコルを通じてのみアクセス可能かどうか
  //     secure: false, // HTTPS接続時のみクッキーを送信するかどうか
  //     // sameSite: 'strict', // SameSite属性
  //     domain: 'localhost', // クッキーの有効なドメイン
  //   });
  // context.res.setHeader('Set-Cookie', serializedCookie);

  return {
    props: {
      user,
    },
  };
};

export default LoginPage;
