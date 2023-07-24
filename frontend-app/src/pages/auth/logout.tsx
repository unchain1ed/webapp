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
  Typography,
} from "@mui/material";
import { GetServerSideProps, NextPage } from "next";
import axios from "axios";

import { Layout as AuthLayout } from "../../layouts/auth/layout";
import React from "react";

type User = {
  UserId: string;
};

type HomeProps = {
  user: User;
};

const Logout: NextPage<HomeProps> & { getLayout: (page: React.ReactNode) => React.ReactNode } = ({
  user,
}) => {
  const router = useRouter();
  const [userId, setUserId] = useState("id");

  useEffect(() => {
    const fetchData = async () => {
      const hostname = process.env.NODE_ENV === "production" ? "server-app" : "localhost";
      try {
        const response = await axios.get(`http://${hostname}:8080/api/login-id`, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          withCredentials: true,
        });
        setUserId(response.data.id);
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
        { userId },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
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
              <Typography variant="h4">Logout</Typography>
              <Typography color="text.secondary" variant="body2"></Typography>
            </Stack>
            <Tab label="Acoount" />
              <Stack spacing={3}>
                <Box sx={{ p: 2, backgroundColor: "#f5f5f5", borderRadius: 4 }}>
                  <Typography variant="caption">ID</Typography>
                  <Typography variant="h5">{userId}</Typography>
                </Box>
              </Stack>
              <Button
                fullWidth
                size="large"
                sx={{ mt: 3 }}
                variant="contained"
                value={userId}
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

Logout.getLayout = (page: React.ReactNode) => <AuthLayout>{page}</AuthLayout>;

export default Logout;
