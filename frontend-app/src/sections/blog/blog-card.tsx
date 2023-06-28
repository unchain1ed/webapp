import { Avatar, Box, Card, CardContent, Divider, Stack, SvgIcon, Typography } from '@mui/material';
import { format } from 'date-fns';

import ArrowDownOnSquareIcon from '@heroicons/react/24/solid/ArrowDownOnSquareIcon';
import ClockIcon from '@heroicons/react/24/solid/ClockIcon';

type Blog = {
  ID: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
};

type BlogCardProps = {
  blog: Blog;
  onClick: (id: string) => void;
};


export const BlogCard: React.FC<BlogCardProps> = ({ blog, onClick }) => {
  const handleClick = () => {
    onClick(blog.ID);
  };

  return (
    <div onClick={handleClick}>
    <Card
      sx={{
        display: 'flex',
        flexDirection: 'column',
        height: '100%'
      }}
    >
      <CardContent>
        <Typography
          align="center"
          gutterBottom
          variant="h5"
        >
          {blog.title}
        </Typography>
        <Typography
          align="center"
          variant="body1"
        >
          {blog.content.split('\n')[0]}
        </Typography>
      </CardContent>
      <Box sx={{ flexGrow: 1 }} />
      <Divider />
      <Stack
        alignItems="center"
        direction="row"
        justifyContent="space-between"
        spacing={2}
        sx={{ p: 2 }}
      >
        <Stack
          alignItems="center"
          direction="row"
          spacing={1}
        >
          <SvgIcon
            color="action"
            fontSize="small"
          >
            <ClockIcon />
          </SvgIcon>
          <Typography
            color="text.secondary"
            display="inline"
            variant="body2"
          >
            {format(new Date(blog.createdAt), 'yyyy/MM/dd HH:mm')}
          </Typography>
        </Stack>
        <Stack
          alignItems="center"
          direction="row"
          spacing={1}
        >
          <SvgIcon
            color="action"
            fontSize="small"
          >
            <ArrowDownOnSquareIcon />
          </SvgIcon>
          <Typography
            color="text.secondary"
            display="inline"
            variant="body2"
          >
            {blog.downloads} Good
          </Typography>
        </Stack>
      </Stack>
    </Card>
    </div>
  );
};

export default BlogCard;
