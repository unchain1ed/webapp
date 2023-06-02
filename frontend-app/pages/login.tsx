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

  const response = await axios.get("http://localhost:8080/login", {
      withCredentials: true,
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });
  const user = response.data;

  return {
    props: {
      user,
    },
  };
};

export default LoginPage;
