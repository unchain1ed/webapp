import axios from "axios";
import { GetServerSideProps, NextPage } from "next";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

type User = {
  LoggedIn: boolean;
  UserId: string;
};

type HomeProps = {
  user: User;
};

const LoginPage: NextPage<HomeProps> = ({ user }) => {
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();
  const [loggedIn, setLoggedIn] = useState(user.LoggedIn);

  useEffect(() => {
    // userの変更を監視する
    if (user.LoggedIn) {
      // ログイン済みの場合の処理
      console.log("ユーザーはログイン済みです");
    } else {
      // ログインしていない場合の処理
      console.log("ユーザーはログインしていません");
    }

    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8080/login", {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        });
        const data = response.data;
        // レスポンスデータの処理
      } catch (error) {
        console.error(error);
      }
    };

    // コンポーネントのマウント時にリクエストを実行
    fetchData();
  }, [user]);

  const handleSubmit = (event: React.MouseEvent<HTMLElement>) => {
    axios
      .post(
        "http://localhost:8080/login",
        { userId, password },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        }
      )
      .then(() => {
        // ログイン成功時の処理
        // window.location.hrefを使用してリダイレクト
        window.location.href = "/mypage";
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
  const response = await axios.get("http://localhost:8080/login", {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  });
  const user = response.data;
console.log(response.data)
// console.log(response)
  return {
    props: {
      user,
    },
  };
};

export default LoginPage;
