import React, { useState, ChangeEvent, FormEvent, useEffect } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import * as Yup from "yup";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { useFormik } from "formik";
import axios from "axios";
import NextLink from "next/link";

import { Logo } from "../../components/logo";

interface BlogForm {
  LoginID: string;
  Title: string;
  Content: string;
}

const Post: React.FC = () => {
  const router = useRouter();

  const [blogForm, setBlogForm] = useState<BlogForm>({
    LoginID: "",
    Title: "",
    Content: "",
  });

  useEffect(() => {
    const fetchData = async () => {
      const hostname = process.env.NODE_ENV === 'production' ? 'server-app' : 'localhost';
      try {
        const response = await axios.get(`http://${hostname}:8080/api/login-id`, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        });
        const fetchedID = response.data.id || null;
        setBlogForm((prevBlogForm) => ({
          ...prevBlogForm,
          LoginID: fetchedID,
        }));
      } catch (error) {
        console.error(error);
      }
    };
    // コンポーネントのマウント時にリクエストを実行
    fetchData();
  }, []);

  const handleSubmit = (e: FormEvent) => {
    // ブログ投稿後にブログ一覧ページにリダイレクト
    router.push("/blog/overview");
  };

  const handlePost = async (event: React.MouseEvent<HTMLElement>) => {
    const hostname = process.env.NODE_ENV === 'production' ? 'server-app' : 'localhost';
    try {
      const response = await axios.post(
        `http://${hostname}:8080/blog/post`,
        blogForm, // blogFormオブジェクトを直接送信
        {
          headers: {
            "Content-Type": "application/json", // JSON形式で送信するためのヘッダー設定
          },
          withCredentials: true,
        }
      );

      router.push("/blog/overview");
      // ログイン成功時の追加の処理を追記する場合はここに記述する
    } catch (error) {
      // ログイン失敗時の処理
      console.error(error);
      // ログイン失敗時の追加の処理を追記する場合はここに記述する
    }
  };

  const formik = useFormik({
    initialValues: {
      title: "",
      content: "",
    },
    validationSchema: Yup.object({
      title: Yup.string().max(50).required("タイトルを入力してください"),
      content: Yup.string().max(10000).required("記事内容を入力してください"),
    }),
    onSubmit: async (values, helpers) => {
      // フォームの送信時の処理
      // 例: APIリクエストなど
      // helpers.setSubmitting(false);  // 必要に応じてフォームを再度利用可能にする
    },
  });

  return (
    <>
      <Head>
        <title>Blog</title>
      </Head>
      <Box sx={{ p: 3 }}>
        <Box
          component={NextLink}
          href="/blog/overview"
          sx={{
            display: "inline-flex",
            height: 32,
            width: 32,
          }}
        >
          <Logo />
        </Box>
      </Box>
      <Box
      >
        <Grid container spacing={2}>
          <Grid item xs={12} md={6}>
            <Box sx={{ width: "100%" }}>
              <Typography variant="h1" component="h1" gutterBottom>
                Create
              </Typography>
              <form onSubmit={handleSubmit}>
                <Stack spacing={2}>
                  <TextField
                    error={!!(formik.touched.title && formik.errors.title)}
                    helperText={formik.touched.title && formik.errors.title}
                    id="title"
                    label="Title"
                    onBlur={formik.handleBlur}
                    value={blogForm.Title && formik.values.title}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setBlogForm({ ...blogForm, Title: event.target.value });
                    }}
                    fullWidth
                  />
                  <TextField
                    error={!!(formik.touched.content && formik.errors.content)}
                    helperText={formik.touched.content && formik.errors.content}
                    id="content"
                    label="Content"
                    onBlur={formik.handleBlur}
                    value={blogForm.Content && formik.values.content}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setBlogForm({ ...blogForm, Content: event.target.value });
                    }}
                    multiline
                    rows={15}
                    fullWidth
                  />
                </Stack>
                <Button
                  size="medium"
                  sx={{ mt: 3 }}
                  type="submit"
                  variant="contained"
                  onClick={handlePost}
                  disabled={!formik.isValid}
                >
                  Submit
                </Button>
              </form>
            </Box>
          </Grid>
          <Grid item xs={12} md={6}>
            <Typography variant="h2" component="h2" gutterBottom>
              Preview
            </Typography>
            {formik.values.title && (
              <Box
                sx={{
                  width: "100%",
                  backgroundColor: "#d4d4d4",
                  padding: "16px",
                  borderRadius: "4px",
                }}
              >
                <Typography variant="body1" style={{ marginTop: "8px", whiteSpace: "pre-wrap" }}>
                  {formik.values.title}
                </Typography>
              </Box>
            )}
            {formik.values.content && (
              <Box
                sx={{
                  width: "100%",
                  backgroundColor: "#f5f5f1",
                  padding: "16px",
                  borderRadius: "4px",
                }}
              >
                <Typography variant="body1" style={{ whiteSpace: "pre-wrap" }}>
                  {formik.values.content}
                </Typography>
              </Box>
            )}
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

// Post.getLayout = (page) => <BlogLayout>{page}</BlogLayout>;

export default Post;
