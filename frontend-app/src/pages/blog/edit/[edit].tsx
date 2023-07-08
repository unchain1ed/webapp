import React, { useState, ChangeEvent, FormEvent, useEffect } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import * as Yup from "yup";

import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { useFormik } from "formik";
import axios from "axios";
import { GetServerSideProps } from "next";

type Blog = {
  ID: string;
  LoginID: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
};

type BlogProps = {
  blog: Blog;
};

const Edit: React.FC = (props: any, context: any) => {
  const { blog } = props;
  const router = useRouter();

  const formik = useFormik({
    initialValues: {
      id: "id",
      title: "title",
      content: "content",
    },
    validationSchema: Yup.object({
      title: Yup.string().max(50).required("タイトルを入力してください"),
      content: Yup.string().max(10000).required("記事内容を入力してください"),
    }),
    onSubmit: async (values, helpers) => {},
  });

  useEffect(() => {
    if (blog) {
      formik.setFieldValue("id", String(blog.ID)); //数値型からstringに変換
      formik.setFieldValue("title", blog.Title);
      formik.setFieldValue("content", blog.Content);
    }
  }, [blog]);

  useEffect(() => {
    console.log("通過E" + formik.values.title);
    console.log("通過R" + formik.values);
  }, []);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    // ブログ投稿後にブログ一覧ページにリダイレクト
    router.push("/blog/overview");
  };

  const handlePost = async (event: React.MouseEvent<HTMLElement>) => {
    try {
      const response = await axios.post("http://localhost:8080/blog/edit", formik.values, {
        headers: {
          "Content-Type": "application/json", // JSON形式で送信するためのヘッダー設定
        },
        withCredentials: true,
      });

      router.push("/blog/overview");
      // ログイン成功時の追加の処理を追記する場合はここに記述する
    } catch (error) {
      // ログイン失敗時の処理
      console.error(error);
      // ログイン失敗時の追加の処理を追記する場合はここに記述する
    }
  };

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
                Edit
              </Typography>
              <form onSubmit={handleSubmit}>
                <Stack spacing={2}>
                  <TextField
                    error={!!(formik.touched.title && formik.errors.title)}
                    helperText={formik.touched.title && formik.errors.title}
                    id="title"
                    label="Title"
                    onBlur={formik.handleBlur}
                    value={formik.values.title}
                    onChange={formik.handleChange}
                    fullWidth
                  />
                  <TextField
                    error={!!(formik.touched.content && formik.errors.content)}
                    helperText={formik.touched.content && formik.errors.content}
                    id="content"
                    label="Content"
                    onBlur={formik.handleBlur}
                    value={formik.values.content}
                    onChange={formik.handleChange}
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

export const getServerSideProps: GetServerSideProps<BlogProps> = async (context) => {
  const { edit } = context.params; // URLからedit(id)を取得

  const response = await axios.get(`http://localhost:8080/blog/overview/post/${edit}`, {
    headers: {
      "Content-Type": "aapplication/json",
    },
    withCredentials: true,
  });

  const blog: Blog = response.data.blog;
  console.log(blog);
  return {
    props: {
      blog,
    },
  };
};

export default Edit;
