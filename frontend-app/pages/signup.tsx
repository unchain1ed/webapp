import axios from "axios";
import { NextPage } from "next";
import router from "next/router";
import { useState } from "react";

const SignupPage: NextPage = () => {
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (event: React.MouseEvent<HTMLElement>) => {
    axios.post("http://localhost:8080/signup", { userId, password }
    )
    .then(() => {
      // サインアップ成功時の処理
      window.location.href = "/"; // リダイレクト先のURLを指定
    })
    .catch((error) => {
      // ログイン失敗時の処理
      console.error(error);
    });
    // try {
    //   axios.post("http://localhost:8080/signup", { userId, password });
    //   // 登録成功時の処理
    //   window.location.href = "/"; // リダイレクト先のURLを指定
    // } catch (error) {
    //   // 登録失敗時の処理
    //   console.error(error);
    // }
  };

  return (
    <>
      <h1>SIGNUP</h1>
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
            新規会員登録
            </button>
          <button type="button" onClick={() => router.push("/")}>
            戻る
          </button>
        </p>
      </form>
    </>
  );
};

export default SignupPage;
