import { useCallback, useEffect, useState } from "react";
import Head from "next/head";
import NextLink from "next/link";
import { useRouter } from "next/router";
import * as Yup from "yup";
import {
  Box,
  Button,
  Stack,
  Tab,
  TextField,
  Typography,
} from "@mui/material";
import { GetServerSideProps, NextPage } from "next";
import axios from "axios";

import { Layout as AuthLayout } from "../../layouts/auth/layout";
import React from "react";
import Layout from "../../layouts/dashboard/layout";

type User = {
  userId: string;
  password: string;
};

type HomeProps = {
  user: User;
};

interface BlogForm {
  loginID: string;
  title: string;
  content: string;
}

const Logout: NextPage<HomeProps> & { getLayout: (page: React.ReactNode) => React.ReactNode } = ({
  user,
}) => {
  const router = useRouter();
  // const [userId, setUserId] = useState("id");
  const [userForm, setUserForm] = useState<User>({
    userId: "",
    password: "",
  });

  useEffect(() => {
    const fetchData = async () => {
      const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
      try {
        const response = await axios.get(`http://${hostname}:8080/api/login-id`, {
          withCredentials: true,
        });
        // setUserId(response.data.id);
        setUserForm((prevBlogForm) => ({
          ...prevBlogForm,
          userId: response.data.id,
        }));
      } catch (error) {
        console.error(error);
      }
    };
    // コンポーネントのマウント時にリクエストを実行
    fetchData();
  }, []);

  const handleLogout = async (event: React.MouseEvent<HTMLElement>) => {
    const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
    try {
      const response = await axios.post(
        `http://${hostname}:8080/logout`,
        userForm,
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );
      router.push("/auth/login");
    } catch (error) {
      // ログイン失敗時の処理
      console.error(error);
    }
  };

  return (
    <>
      <Head>
        <title>Logout</title>
      </Head>
      <Box
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
              <Typography variant="h4">Logout</Typography>
              <Typography color="text.secondary" variant="body2"></Typography>
            </Stack>
            <Tab label="Acoount" />
              <Stack spacing={3}>
                <Box sx={{ p: 2, backgroundColor: "#f5f5f5", borderRadius: 4 }}>
                  <Typography variant="caption">ID</Typography>
                  <Typography variant="h5">{userForm.userId}</Typography>
                </Box>
              </Stack>
              <Stack spacing={3} sx={{ mt: 4 }}>
            <TextField
              required
              type="password"
              label="Password"
              value={userForm.password}
              onChange={(event) => {
                setUserForm({ ...userForm, password: event.target.value });
              }}
              fullWidth
            />
          </Stack>
              <Button
                size="medium"
                sx={{ mt: 3 }}
                variant="contained"
                value={userForm.userId}
                onClick={handleLogout}
              >
                Logout
              </Button>
          </div>
        </Box>
      </Box>
    </>
  );
};

// export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
//   const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
//   const response = await axios.get(`http://${hostname}:8080/api/login-id`, {
//     headers: {
//       "Content-Type": "application/x-www-form-urlencoded",
//     },
//     withCredentials: true,
//   });
//   const user = response.data;
//   return {
//     props: {
//       user,
//     },
//   };
// };

// Logout.getLayout = (page: React.ReactNode) => <AuthLayout>{page}</AuthLayout>;
Logout.getLayout = (page: React.ReactNode) => (
  <Layout>
    {page}
  </Layout>
);
export default Logout;
