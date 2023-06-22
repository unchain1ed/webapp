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
import { Layout as DashboardLayout } from 'src/layouts/dashboard/layout';
import { BlogCard } from 'src/sections/blog/blog-card';
import { CompaniesSearch } from 'src/sections/blog/companies-search';
import axios from 'axios';
import router from 'next/router';
import React from 'react';
import { GetServerSideProps } from 'next';

type Blog = {
  ID: string;
  title: string;
  content: string;
  updatedAt: string;
};

type BlogProps = {
  blogs: Blog[];
};

const handleAdd = (event: React.MouseEvent<HTMLElement>) => {
  axios
    .get(
      "http://localhost:8080/",
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

// const companies = [
//   {
//     id: '2569ce0d517a7f06d3ea1f24',
//     createdAt: '27/03/2019',
//     description: 'Dropbox is a file hosting service that offers cloud storage, file synchronization, a personal cloud.',
//     logo: '/assets/logos/logo-dropbox.png',
//     title: 'Dropbox',
//     downloads: '594'
//   }
// ];

const Page = ({ blogs }: BlogProps) => (
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
                <Button
                  color="inherit"
                  startIcon={(
                    <SvgIcon fontSize="small">
                      <ArrowUpOnSquareIcon />
                    </SvgIcon>
                  )}
                >
                  Import
                </Button>
                <Button
                  color="inherit"
                  startIcon={(
                    <SvgIcon fontSize="small">
                      <ArrowDownOnSquareIcon />
                    </SvgIcon>
                  )}
                >
                  Export
                </Button>
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
                Add
              </Button>
            </div>
          </Stack>
          <CompaniesSearch />
          <Grid
            container
            spacing={3}
          >
            {blogs.map((blogs) => (
              <Grid
                xs={12}
                md={6}
                lg={4}
                key={blogs.ID}
              >
                <BlogCard blog={blogs} /> 
                {/* ↑TODOここ */}
              </Grid>
            ))}
          </Grid>
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center'
            }}
          >
            <Pagination
              count={3}
              size="small"
            />
          </Box>
        </Stack>
      </Container>
    </Box>
  </>
);

Page.getLayout = (page: any) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);


export const getServerSideProps: GetServerSideProps<BlogProps> = async (context) => {
  
  const response = await axios.get("http://localhost:8080/blog/overview", {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    withCredentials: true,
  });

  const blogsInfo = response.data.blogs;

  const blogs: Blog[] = blogsInfo.map((item: any) => ({
    ID: item.ID,
    title: item.Title,
    content: item.Content,
    updatedAt: item.UpdatedAt
    }));

  return {
    props: {
      blogs,
    },
  };
};

export default Page;
