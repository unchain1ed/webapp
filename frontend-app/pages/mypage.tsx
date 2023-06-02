import axios from "axios";
import { GetServerSideProps, NextPage } from "next";
import { useState } from "react";


type User = {
  ID: number;
  UserId: string;
};

type MyPageProps = {
  user: User;
};

const MyPage: NextPage<MyPageProps> = ({ user }) => {

    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [error, setError] = useState('');
  
    const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
      setTitle(event.target.value);
    };
  
    const handleContentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
      setContent(event.target.value);
    };
  
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
  
      try {
        const response = await axios.post('/api/posts', { title, content });
        // Handle successful post creation (e.g., show success message, redirect, etc.)
      } catch (error) {
        // setError(response.data.message);
      }
    };

  const handleEditProfile = () => {
    // プロフィール編集画面への遷移
    window.location.href = "/edit-profile";
  };

  const handleLogout = () => {
    // ログアウトAPIを呼び出してセッションを破棄する
    axios.get("/logout")
      .then(() => {
        // ログアウト成功時の処理
        window.location.href = "/"; // ホーム画面への遷移
      })
      .catch(error => {
        // ログアウト失敗時の処理
        console.error(error);
      });
  };

  return (
    <>
      <h1>MYPAGE</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="title">Title:</label>
          <input type="text" id="title" value={title} onChange={handleTitleChange} required />
        </div>
        <div>
          <label htmlFor="content">Content:</label>
          <textarea id="content" value={content} onChange={handleContentChange} required />
        </div>
        {error && <p>Error: {error}</p>}
        <button type="submit">Create Post</button>
      </form>

      {user && (
        <>
          <p>ログイン中 {user.ID}【{user.UserId}】</p>
          <p>ログイン成功</p>
          <button onClick={handleEditProfile}>プロフィール編集</button>
          <button onClick={handleLogout}>ログアウト</button>
        </>
      )}
      <a href="/">ホーム</a>
      <a href="/signup">会員登録</a>
      <a href="/login">ログイン</a>
    </>
  );
};

// export const getServerSideProps: GetServerSideProps<MyPageProps> = async () => {
//   //TLS/SSL接続時に証明書を検証せずに接続を許可するかどうかを制御 "0" 検証を無効
//   // process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

//   const response = await axios.get("http://localhost:8080/mypage", {
//       // headers: {
//       //   "Content-Type": "application/x-www-form-urlencoded",
//       // },
//       withCredentials: true,
//     });
//   const user = response.data;

//   return {
//     props: {
//       user,
//     },
//   };
// };

export default MyPage;
