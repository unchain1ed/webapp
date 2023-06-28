import React, { useState, ChangeEvent, FormEvent } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import * as Yup from "yup";

import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { useFormik } from "formik";
import axios from "axios";

interface BlogForm {
  title: string;
  content: string;
}

const Post: React.FC = () => {
  const router = useRouter();

  const [blogForm, setBlogForm] = useState<BlogForm>({
    title: "",
    content: "",
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    // ブログ投稿の処理
    // APIリクエストなどが含まれます

    // ブログ投稿後にブログ一覧ページにリダイレクト
    router.push("/blog/overview");
  };

  const handlePost = async (event: React.MouseEvent<HTMLElement>) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/blog/post",
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
      <Box
        component="main"
        sx={{
          display: "flex",
          justifyContent: "center",
          py: 8,
        }}
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
                    value={blogForm.title && formik.values.title}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setBlogForm({ ...blogForm, title: event.target.value });
                    }}
                    fullWidth
                  />
                  <TextField
                    error={!!(formik.touched.content && formik.errors.content)}
                    helperText={formik.touched.content && formik.errors.content}
                    id="content"
                    label="Content"
                    onBlur={formik.handleBlur}
                    value={blogForm.content && formik.values.content}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setBlogForm({ ...blogForm, content: event.target.value });
                    }}
                    multiline
                    rows={15}
                    fullWidth
                  />
                </Stack>
                <Button
                  fullWidth
                  size="large"
                  sx={{ mt: 3 }}
                  type="submit"
                  variant="contained"
                  onClick={handlePost}
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
                <Typography variant="pre" style={{ marginTop: "8px", whiteSpace: "pre-wrap" }}>
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
                <Typography variant="pre" style={{ whiteSpace: "pre-wrap" }}>
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
