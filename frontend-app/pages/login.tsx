import axios from "axios";
import { NextPage } from "next";
import { useState } from "react";


const LoginPage: NextPage = () => {
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    axios.post("/login", { userId, password }, {
        method: 'POST',
        // mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
        }
      })
    //   .then(() => {
    //     // ログイン成功時の処理
    //     window.location.href = "/mypage"; // リダイレクト先のURLを指定
    //   })
    //   .catch((error) => {
    //     // ログイン失敗時の処理
    //     console.error(error);
    //   });
  };

  return (
    <>
      <h1>LOGIN</h1>
      <form onSubmit={handleSubmit}>
        <p>
          ユーザーID：
          <input
            type="text"
            name="user_id"
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
          <input type="submit" value="ログイン" />
          <a href="/">
            <input type="button" value="戻る" />
          </a>
        </p>
      </form>
    </>
  );
};

export default LoginPage;
