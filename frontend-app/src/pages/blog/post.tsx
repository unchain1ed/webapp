import React, { useState, ChangeEvent, FormEvent } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import * as Yup from "yup";
import { Layout as BlogLayout } from "src/layouts/blog/layout";

import { Box, Button, Grid, Stack, TextField } from "@mui/material";
import { useFormik } from "formik";

const Page: React.FC = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const router = useRouter();

  const handleTitleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleContentChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setContent(e.target.value);
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    // ブログ投稿の処理
    // APIリクエストなどが含まれます

    // ブログ投稿後にブログ一覧ページにリダイレクト
    router.push("/blogs");
  };

  const formik = useFormik({
    initialValues: {
      // userId: "root",
      // password: "root",
      // submit: null,
    },
    validationSchema: Yup.object({
      title: Yup.string().max(30).required("タイトルを入力してください"),
      content: Yup.string().max(20).required("記事内容を入力してください"),
    }),
    onSubmit: async (values, helpers) => {},
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
              <h1>Create</h1>
              <form onSubmit={handleSubmit}>
                <Stack spacing={2}>
                  <TextField
                    error={!!(formik.touched.title && formik.errors.title)}
                    helperText={formik.touched.title && formik.errors.title}
                    id="title"
                    label="Title"
                    onBlur={formik.handleBlur}
                    value={title && formik.values.title}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setTitle(event.target.value);
                    }}
                    fullWidth
                  />
                  <TextField
                    error={!!(formik.touched.content && formik.errors.content)}
                    helperText={formik.touched.content && formik.errors.content}
                    id="content"
                    label="Content"
                    onBlur={formik.handleBlur}
                    value={content && formik.values.content}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setContent(event.target.value);
                    }}
                    multiline
                    rows={15}
                    fullWidth
                  />
                </Stack>
                <Button type="submit" variant="contained">
                  Submit
                </Button>
              </form>
            </Box>
          </Grid>
          <Grid item xs={12} md={6}>
            <Box
              sx={{
                width: "100%",
                backgroundColor: "#f5f5f5",
                padding: "16px",
                borderRadius: "4px",
              }}
            >
              <h2>Preview</h2>
              <pre style={{ marginTop: "8px", whiteSpace: "pre-wrap" }}>{title}</pre>
              <pre style={{ whiteSpace: "pre-wrap" }}>{content}</pre>
            </Box>
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

// Page.getLayout = (page) => <BlogLayout>{page}</BlogLayout>;

export default Page;
