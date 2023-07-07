import React, { useState, ChangeEvent, FormEvent, useEffect } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import * as Yup from "yup";

import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { useFormik } from "formik";
import axios from "axios";
import { GetServerSideProps } from "next";

interface BlogForm {
  // ID: string;
  LoginID: string;
  title: string;
  content: string;
}
type Blog = {
  // ID: string;
  LoginID: string;
  title: string;
  content: string;
  // createdAt: Date;
  // updatedAt: Date;
  // deletedAt: Date;
};

type BlogProps = {
  blog: Blog;
};

const initialValues = {
  title: "",
  content: "",
};

const Edit: React.FC = (props: any, context: any) => {

  const { blog } = props;

  // const [blogForm, setBlogForm] = useState<BlogForm>({
  //   ID: "",
  //   LoginID: "",
  //   title: "",
  //   content: "",
  // });

  // const [blogForm, setBlogForm] = useState<BlogForm>({
  //   // ID: "",
  //   LoginID: blog.LoginID,
  //   title: blog.title,
  //   content: blog.content,
  //   // createdAt: new Date,
  //   // updatedAt: new Date,
  //   // deletedAt: new Date,
  // });

  const formik = useFormik({
    initialValues: {
      title: "blog.title",
      content: "blog.content",
    },
    validationSchema: Yup.object({
      title: Yup.string().max(50).required("タイトルを入力してください"),
      content: Yup.string().max(10000).required("記事内容を入力してください"),
    }),
    onSubmit: async (values, helpers) => {
      // フォームの送信時の処理
      // helpers.setSubmitting(false);  // 必要に応じてフォームを再度利用可能にする
    },
  });

  // useEffect(() => {
  //   setBlogForm({
  //     LoginID: blog.LoginID,
  //     title: blog.title,
  //     content: blog.content,
  //   });
  // }, [blog]);

  // useEffect(() => {

  //   const fetchData = () => {
  //   setBlogForm({
  //     LoginID: blog.LoginID,
  //     title: blog.title,
  //     content: blog.content,
  //   })};
  //   // コンポーネントのマウント時にリクエストを実行
  //   fetchData();
  // }, []);
  
  useEffect(() => {
        if (blog) {
          formik.setFieldValue("title", blog.Title);
          formik.setFieldValue("content", blog.Content);
        }
  }, [blog]);

  // useEffect(() => {
  //   if (blog) {
  //     formik.handleChange({ target: { name: "title", value: blog.title } });
  //     formik.handleChange({ target: { name: "content", value: blog.content } });
  //   }
  //   console.log("通過E" + formik.values.title);
  // }, [blog]);

  // useEffect(() => {
  //   // サーバーから前回の入力内容を取得する非同期処理
  //   const fetchData = async () => {
  //     try {
  //       // const { edit } = context.params; // URLからedit(id)を取得
  //       console.log("edit"+ router.query.edit)
  //       const response = await axios.get(`http://localhost:8080/blog/overview/post/${router.query.edit}`, {
  //         headers: {
  //           "Content-Type": "application/x-www-form-urlencoded",
  //         },
  //         withCredentials: true,
  //       });
  //       const blog = response.data.blog;
  //       console.log("通過B", JSON.stringify(blog));
  //       // サーバーから取得した前回の入力内容を設定
  //       formik.setValues({
  //         title: blog.title,
  //         content: blog.content,
  //       });
  //     } catch (error) {
  //       console.error("Error fetching data:", error);
  //     }
  //   };

  //   // fetchData(context);
  // }, []);

  // useEffect(() => {
  //   const fetchData = async () => {
  //     try {
  //       const response = await axios.get(`http://localhost:8080/blog/overview/post/${router.query.edit}`, {
  //         headers: {
  //           "Content-Type": "application/json",
  //         },
  //         withCredentials: true,
  //       });
  //       // console.log("response.data.blog", response.data.blog);
        
  //       const blog = response.data.blog;
  //       // console.log("通過B", JSON.stringify(blog));
  //       // console.log("通過C", blog.title);
  //       // console.log("blog.title", JSON.stringify(blog.Title));
  //       // console.log("blog.content", JSON.stringify(blog.Content));

  //       if (blog) {
  //         formik.setFieldValue("title", blog.Title);
  //         formik.setFieldValue("content", blog.Content);
  //       }
  //     } catch (error) {
  //       console.error("Error fetching data:", error);
  //     }
  //   };
  
  //   fetchData(); // fetchData関数を呼び出す
  // }, []); 
  



  const handleSubmit = (e: FormEvent) => {
    // e.preventDefault();
    // ブログ投稿後にブログ一覧ページにリダイレクト
    // router.push("/blog/overview");
  };

  const handlePost = async (event: React.MouseEvent<HTMLElement>) => {
    try {
      // const response = await axios.post(
      //   "http://localhost:8080/blog/post",
      //   blogForm, // blogFormオブジェクトを直接送信
      //   {
      //     headers: {
      //       "Content-Type": "application/json", // JSON形式で送信するためのヘッダー設定
      //     },
      //     withCredentials: true,
      //   }
      // );

      // router.push("/blog/overview");
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
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  });

  const blog: Blog = response.data.blog;

  return {
    props: {
      blog,
    },
  };
};


export default Edit;
