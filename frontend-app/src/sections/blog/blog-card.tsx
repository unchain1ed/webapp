import { useState } from 'react';
import {
  Avatar,
  Box,
  Card,
  CardContent,
  Divider,
  MenuItem,
  Select,
  Stack,
  SvgIcon,
  Typography,
} from '@mui/material';
import { format } from 'date-fns';
import { SelectChangeEvent } from '@mui/material/SelectChangeEvent'; // 追加

import ArrowDownOnSquareIcon from '@heroicons/react/24/solid/ArrowDownOnSquareIcon';
import ClockIcon from '@heroicons/react/24/solid/ClockIcon';
import router from 'next/router';

type Blog = {
  ID: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
};

type BlogCardProps = {
  blog: Blog;
  onClick: (id: string, action: string) => void;
};

export const BlogCard: React.FC<BlogCardProps> = ({ blog, onClick }) => {
  const [selectedAction, setSelectedAction] = useState('');


  const handleActionChange = (event: SelectChangeEvent<string>) => {
    // const selectedValue = event.target.value as string;
    setSelectedAction(event.target.value as string);
  
    if (selectedAction === "edit") {
      // 編集の処理
      console.log("編集が選択されました");

      router.push(`/blog/edit`); 

    
    } else if (selectedAction === "delete") {
      // 削除の処理
      console.log("削除が選択されました");
    }
  };
  

  const handleClick = () => {
    onClick(blog.ID, selectedAction);
    // router.push(`/blog/post`); 
  };


  const handleContainerClick = (event: React.MouseEvent<HTMLDivElement>) => {
    if (selectedAction != "edit") {
    
    
    // プルダウン以外の部分がクリックされた場合の処理
    console.log("igaino選択されました");
    event.stopPropagation();
    // ここで必要な処理を実行する
    }
  };

  return (
    <div onClick={handleContainerClick}>
      <Card
        sx={{
          display: 'flex',
          flexDirection: 'column',
          height: '100%',
        }}
      >
        <CardContent>
          <Typography align="center" gutterBottom variant="h5">
            {blog.title}
          </Typography>
          <Typography align="center" variant="body1">
            {blog.content.split('\n')[0]}
          </Typography>
        </CardContent>
        <Box sx={{ flexGrow: 1 }} />
        <Divider />
        <Stack alignItems="center" direction="row" justifyContent="space-between" spacing={2} sx={{ p: 2 }}>
          <Stack alignItems="center" direction="row" spacing={1}>
            <SvgIcon color="action" fontSize="small">
              <ClockIcon />
            </SvgIcon>
            <Typography color="text.secondary" display="inline" variant="body2">
              {format(new Date(blog.createdAt), 'yyyy/MM/dd HH:mm')}
            </Typography>
          </Stack>
          <Stack alignItems="center" direction="row" spacing={1}>
            <SvgIcon color="action" fontSize="small">
              <ArrowDownOnSquareIcon />
            </SvgIcon>
            <Typography color="text.secondary" display="inline" variant="body2">
              {blog.downloads} Good
            </Typography>
          </Stack>
          <Select value={selectedAction} onChange={handleActionChange} variant="outlined" size="small">
            <MenuItem value="edit">Edit</MenuItem>
            <MenuItem value="delete">Delete</MenuItem>
          </Select>
        </Stack>
      </Card>
    </div>
  );
};

export default BlogCard;
