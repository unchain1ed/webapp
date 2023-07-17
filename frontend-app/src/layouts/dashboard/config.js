import ChartBarIcon from '@heroicons/react/24/solid/ChartBarIcon';
import CogIcon from '@heroicons/react/24/solid/CogIcon';
import LockClosedIcon from '@heroicons/react/24/solid/LockClosedIcon';
import ShoppingBagIcon from '@heroicons/react/24/solid/ShoppingBagIcon';
import UserIcon from '@heroicons/react/24/solid/UserIcon';
import UserPlusIcon from '@heroicons/react/24/solid/UserPlusIcon';
import UsersIcon from '@heroicons/react/24/solid/UsersIcon';
import XCircleIcon from '@heroicons/react/24/solid/XCircleIcon';
import { SvgIcon } from '@mui/material';

export const items = [
  {
    title: 'Overview',
    path: '/',
    icon: (
      <SvgIcon fontSize="small">
        <ChartBarIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Blog',
    path: '/blog/overview',
    icon: (
      <SvgIcon fontSize="small">
        <ShoppingBagIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Account',
    path: '/account',
    icon: (
      <SvgIcon fontSize="small">
        <UserIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Settings',
    path: '/settings',
    icon: (
      <SvgIcon fontSize="small">
        <CogIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Register',
    path: '/auth/register',
    icon: (
      <SvgIcon fontSize="small">
        <UserPlusIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Login',
    path: '/auth/login',
    icon: (
      <SvgIcon fontSize="small">
        <LockClosedIcon />
      </SvgIcon>
    )
  },
  {
    title: 'Logout',
    path: '/',
    icon: (
      <SvgIcon fontSize="small">
        <LockClosedIcon />
      </SvgIcon>
    )
  },
];



// export const [items2, setItems] = useState([
//   {
//     title: 'Overview',
//     path: '/',
//     icon: (
//       <SvgIcon fontSize="small">
//         <ChartBarIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Customers',
//     path: '/customers',
//     icon: (
//       <SvgIcon fontSize="small">
//         <UsersIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Blog',
//     path: '/blog/overview',
//     icon: (
//       <SvgIcon fontSize="small">
//         <ShoppingBagIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Account',
//     path: '/account',
//     icon: (
//       <SvgIcon fontSize="small">
//         <UserIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Settings',
//     path: '/settings',
//     icon: (
//       <SvgIcon fontSize="small">
//         <CogIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Register',
//     path: '/auth/register',
//     icon: (
//       <SvgIcon fontSize="small">
//         <UserPlusIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Error',
//     path: '/404',
//     icon: (
//       <SvgIcon fontSize="small">
//         <XCircleIcon />
//       </SvgIcon>
//     )
//   },
//   {
//     title: 'Login',
//     path: '/auth/login',
//     icon: (
//       <SvgIcon fontSize="small">
//         <LockClosedIcon />
//       </SvgIcon>
//     )
//   },
// ]);

// const handleConditionChange = (condition) => {
//   if (condition) {
//     // 条件がtrueの場合、配列の一番最後を消去し、配列に新たに要素を追加する
//     setItems(items.slice(0, items.length - 1));
//     setItems([
//       {
//         title: 'Logout',
//         path: '/',
//         icon: (
//           <SvgIcon fontSize="small">
//             <LockClosedIcon />
//           </SvgIcon>
//         )
//       },
//     ]);
//   } else {
//     // 条件がfalseの場合、配列そのまま

//   }
// };