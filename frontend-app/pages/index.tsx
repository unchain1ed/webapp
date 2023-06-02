import { NextPage, GetServerSideProps } from "next";
import { useEffect } from "react";
import axios from "axios";
import router from "next/router";

// axiosのデフォルトヘッダーにOriginを設定
// axios.defaults.headers.common["Origin"] = "http://localhost:3000/";

type User = {
  LoggedIn: boolean;
  UserId: string;
};

type HomeProps = {
  user: User;
};

const Home: NextPage<HomeProps> = ({ user }) => {
  //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  // process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
  useEffect(() => {
    (async () => {
    await axios.get("http://localhost:8080"),  {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
    }
  });
  }, []);

  return (
    <>
      <h1>TOP</h1>
      {/* {user.LoggedIn && <p>ログイン中【{user.UserId}】</p>} */}
      <p>ログイン中【{user.UserId}】</p>
      <button type="button" onClick={() => router.push("/signup")}>
        会員登録
      </button>
      <button type="button" onClick={handleLogin}>
        ログイン
      </button>
      <button type="button" onClick={handleLogout}>
        ログアウト
      </button>
    </>
  );
};

export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
  //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  // process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  // const response = await axios.get("http://localhost:8080");

  const response = await axios.get("http://localhost:8080",  {
  headers: {
    "Content-Type": "application/x-www-form-urlencoded",
  },
  withCredentials: true,
})
  const user = response.data;

  return {
    props: {
      user,
    },
  };
};

const handleLogout = (event: React.MouseEvent<HTMLElement>) => {

  axios.get("http://localhost:8080/logout", {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  })
  .then(() => {
        // ログイン成功時の処理
        window.location.href = "/login"; // リダイレクト先のURLを指定
      })
      .catch((error) => {
        // ログイン失敗時の処理
        console.error(error);
      });
  };

  const handleLogin = (event: React.MouseEvent<HTMLElement>) => {

    axios.get("http://localhost:8080/login", {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      withCredentials: true,
    })
    .then(() => {
          // ログイン成功時の処理
          window.location.href = "/login"; // リダイレクト先のURLを指定
        })
        .catch((error) => {
          // ログイン失敗時の処理
          console.error(error);
        });
    };

export default Home;
