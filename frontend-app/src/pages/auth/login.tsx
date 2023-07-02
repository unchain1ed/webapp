import { useCallback, useEffect, useState } from "react";
import Head from "next/head";
import NextLink from "next/link";
import { useRouter } from "next/router";
import { useFormik } from "formik";
import * as Yup from "yup";
import {
  Alert,
  Box,
  Button,
  FormHelperText,
  Link,
  Stack,
  Tab,
  Tabs,
  TextField,
  Typography,
} from "@mui/material";
import { GetServerSideProps, NextPage } from "next";
import axios from "axios";

import { useAuth } from "src/hooks/use-auth";
import { Layout as AuthLayout } from "src/layouts/auth/layout";
import React from "react";

type User = {
  UserId: string;
};

type HomeProps = {
  user: User;
};

const Page: NextPage<HomeProps> = ({ user }) => {
  const router = useRouter();
  const auth = useAuth();
  const [method, setMethod] = useState("userId");
  const [userId, setUserId] = useState("root");
  const [password, setPassword] = useState("root");

  useEffect(() => {
    console.log("通過useEffect" + user);
    // userの変更を監視する
    if (user.UserId) {
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

  // const handleLogin = async (event: React.MouseEvent<HTMLElement>) => {
  //   try {
  //     const response = await axios.post(
  //       "http://localhost:8080/login",
  //       { userId, password },
  //       {
  //         headers: {
  //           "Content-Type": "application/x-www-form-urlencoded",
  //         },
  //         withCredentials: true,
  //       }
  //     );
  //     // ログイン成功時の処理
  //     await auth.signIn(userId, password);
  //     router.push("/");
  //     // helpers.setStatus({ success: false });
  //     // helpers.setErrors({ submit: err.message });
  //     // helpers.setSubmitting(false);
  //     // ログイン成功時の追加の処理を追記する場合はここに記述する
  //   } catch (error) {
  //     // ログイン失敗時の処理
  //     console.error(error);
  //     // ログイン失敗時の追加の処理を追記する場合はここに記述する
  //   }
  // };

  const handleLogin = async (event: React.MouseEvent<HTMLElement>) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/login",
        { userId, password },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        }
      );
      
      user.UserId = response.data.user.UserId; // レスポンスデータを取得
  
      // フロントエンドのビューで表示するための状態を作成
      // const [displayData, setDisplayData] = useState("");
  
      // // 取得したデータを状態に設定
      // useEffect(() => {
      //   setDisplayData(responseData);
      // }, []);

      console.log("通過displayData", response.data.user.UserId);

      console.log("通過displayData2", user.UserId);

      
      // ログイン成功時の処理
      await auth.signIn(userId, password);
      router.push("/");
      // ログイン成功時の追加の処理を追記する場合はここに記述する
    } catch (error) {
      // ログイン失敗時の処理
      console.error(error);
      // ログイン失敗時の追加の処理を追記する場合はここに記述する
    }
  };
  
  const formik = useFormik({
    initialValues: {
      userId: "root",
      password: "root",
      submit: null,
    },
    validationSchema: Yup.object({
      userId: Yup.string().max(20).required("IDを入力してください"),
      password: Yup.string().max(20).required("Passwordを入力してください"),
    }),
    onSubmit: async (values, helpers) => {},
  });

  const handleMethodChange = useCallback((event: React.SyntheticEvent, value: string) => {
    setMethod(value);
  }, []);

  const handleSkip = useCallback(() => {
    auth.skip();
    router.push("/");
  }, [auth, router]);

  return (
    <>
      <Head>
        <title>Login</title>
      </Head>
      <Box
        sx={{
          backgroundColor: "background.paper",
          flex: "1 1 auto",
          alignItems: "center",
          display: "flex",
          justifyContent: "center",
        }}
      >
        <Box
          sx={{
            maxWidth: 550,
            px: 3,
            py: "100px",
            width: "100%",
          }}
        >
          <div>
            <Stack spacing={1} sx={{ mb: 3 }}>
              <Typography variant="h4">Login</Typography>
              <Typography color="text.secondary" variant="body2">
                Don&apos;t have an account? &nbsp;
                <Link
                  component={NextLink}
                  href="/auth/register"
                  underline="hover"
                  variant="subtitle2"
                >
                  Register
                </Link>
              </Typography>
            </Stack>
            <Tabs onChange={handleMethodChange} sx={{ mb: 3 }} value={method}>
              <Tab label="Acoount" value="userId" />
            </Tabs>
            {method === "userId" && (
              <form noValidate onSubmit={formik.handleSubmit}>
                <Stack spacing={3}>
                  <TextField
                    error={!!(formik.touched.userId && formik.errors.userId)}
                    fullWidth
                    helperText={formik.touched.userId && formik.errors.userId}
                    label="ID"
                    name="userId"
                    onBlur={formik.handleBlur}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setUserId(event.target.value);
                    }}
                    type="text"
                    value={userId && formik.values.userId}
                  />
                  <TextField
                    error={!!(formik.touched.password && formik.errors.password)}
                    fullWidth
                    helperText={formik.touched.password && formik.errors.password}
                    label="Password"
                    name="password"
                    onBlur={formik.handleBlur}
                    onChange={(event) => {
                      formik.handleChange(event);
                      setPassword(event.target.value);
                    }}
                    type="password"
                    value={password && formik.values.password}
                  />
                </Stack>
                <FormHelperText sx={{ mt: 1 }}>Optionally you can skip.</FormHelperText>
                {formik.errors.submit && (
                  <Typography color="error" sx={{ mt: 3 }} variant="body2">
                    {formik.errors.submit}
                  </Typography>
                )}
                <Button
                  fullWidth
                  size="large"
                  sx={{ mt: 3 }}
                  // type="submit"
                  variant="contained"
                  onClick={handleLogin}
                >
                  Continue
                </Button>
                <Button fullWidth size="large" sx={{ mt: 3 }} onClick={handleSkip}>
                  Skip authentication
                </Button>
                <Alert color="primary" severity="info" sx={{ mt: 3 }}>
                  <div>
                    You can use <b>root</b> and password <b>root</b>
                  </div>
                </Alert>
              </form>
            )}
          </div>
        </Box>
      </Box>
    </>
  );
};

Page.getLayout = (page: React.ReactNode) => <AuthLayout>{page}</AuthLayout>;

export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
  const response = await axios.get("http://localhost:8080/login", {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  });
  const user = response.data;
  // console.log(response.data)
  console.log("通過getServerSideProps" + user.UserId);
  return {
    props: {
      user,
    },
  };
};

export default Page;
