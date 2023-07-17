import Head from 'next/head';
import ArrowUpOnSquareIcon from '@heroicons/react/24/solid/ArrowUpOnSquareIcon';
import ArrowDownOnSquareIcon from '@heroicons/react/24/solid/ArrowDownOnSquareIcon';
import PlusIcon from '@heroicons/react/24/solid/PlusIcon';
import {
  Box,
  Button,
  Container,
  Pagination,
  Stack,
  SvgIcon,
  Typography,
  Unstable_Grid2 as Grid
} from '@mui/material';
import Layout from '../../layouts/dashboard/layout';
import { BlogCard } from '../../sections/blog/blog-card';
import { CompaniesSearch } from '../../sections/blog/companies-search';
import axios from 'axios';
import router from 'next/router';
import React from 'react';
import { GetServerSideProps } from 'next';

type Blog = {
  ID: string;
  LoginID: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
};

type BlogProps = {
  blogs: Blog[];
};

type ClickValue = {
  value: string;
};

const handleAdd = (event: React.MouseEvent<HTMLElement>) => {
  const hostname = process.env.NODE_ENV === 'production' ? 'server-app' : 'localhost';
  axios
    .get(
      `http://${hostname}:8080/`,
      {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        withCredentials: true,
      }
    )
    .then(() => {
      // ログイン成功時の処理
      router.push('/blog/post');
      // helpers.setStatus({ success: false });
      // helpers.setErrors({ submit: err.message });
      // helpers.setSubmitting(false);
    })
    .catch((error) => {
      // ログイン失敗時の処理
      console.error(error);
    });
};

const Overview = ({ blogs }: BlogProps, { value }: ClickValue) => {

  const handleBlogClick = (id: string) => {
  };

  return (
  <>
    <Head>
      <title>
        Blog
      </title>
    </Head>
    <Box
      component="main"
      sx={{
        flexGrow: 1,
        py: 8
      }}
    >
      <Container maxWidth="xl">
        <Stack spacing={3}>
          <Stack
            direction="row"
            justifyContent="space-between"
            spacing={4}
          >
            <Stack spacing={1}>
              <Typography variant="h4">
                Blog
              </Typography>
              <Stack
                alignItems="center"
                direction="row"
                spacing={1}
              >
              </Stack>
            </Stack>
            <div>
              <Button
                startIcon={(
                  <SvgIcon fontSize="small">
                    <PlusIcon />
                  </SvgIcon>
                )}
                variant="contained"
                onClick={handleAdd}
              >
                Create New Blog
              </Button>
            </div>
          </Stack>
          <CompaniesSearch />
          <Grid container spacing={3}>
            {blogs.map((blog) => (
          <Grid xs={12} md={6} lg={4} key={blog.ID}>
            <BlogCard blog={blog} clickValue={value} onClick={handleBlogClick} />
          </Grid>
            ))}
          </Grid>
        </Stack>
      </Container>
    </Box>
  </>
  );
};

Overview.getLayout = (page: any) => (
  <Layout>
    {page}
  </Layout>
);


export const getServerSideProps: GetServerSideProps<BlogProps> = async (context) => {
  const hostname = process.env.NODE_ENV === 'production' ? 'server-app' : 'localhost';
  const response = await axios.get(`http://${hostname}:8080/blog/overview`, 
   {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  });

  const blogsInfo = response.data.blogs;


  const blogs: Blog[] = blogsInfo.map((item: any) => ({
    ID: item.ID,
    LoginID: item.LoginID,
    title: item.Title,
    content: item.Content,
    createdAt: item.CreatedAt,
    updatedAt: item.UpdatedAt
    }));

  return {
    props: {
      blogs,
    },
  };
};

export default Overview;
