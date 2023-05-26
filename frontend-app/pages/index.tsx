import { NextPage, GetServerSideProps } from "next";
import { useEffect } from "react";
import axios from "axios";

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
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
  useEffect(() => {
    (async () => {
    await axios.get("https://localhost:8080");
    })();
  }, []);

  return (
    <>
      <h1>TOP</h1>
      {user.LoggedIn && <p>ログイン中【{user.UserId}】</p>}
      <a href="signup">
        <input type="button" value="会員登録" />
      </a>
      <a href="login">
        <input type="button" value="ログイン" />
      </a>
      <a href="logout">
        <input type="button" value="ログアウト" />
      </a>
    </>
  );
};

export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
  //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
  process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

  const response = await axios.get("https://localhost:8080");
  const user = response.data;

  return {
    props: {
      user,
    },
  };
};

export default Home;
