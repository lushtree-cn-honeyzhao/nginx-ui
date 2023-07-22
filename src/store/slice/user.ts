import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import {User} from "../../models/user.ts";

export type IUserState = {
  user?: User;
  isLogin: boolean;
  isAdmin?: boolean;
};

const initialState: IUserState = {
  isLogin: false,
  isAdmin: false,
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<User>) {
      state.user = action.payload;
      state.isLogin = true;
      const roles = state.user?.roles || '';
      state.isAdmin = roles.indexOf('ADMIN') > -1;
    },
    clearUser(state) {
      state.isLogin = false;
      state.user = undefined;
      state.isAdmin = false;
      console.log('======clearUser======');
    },
  },
});

export const { setUser, clearUser } = userSlice.actions;
export default userSlice.reducer;
export const UserActions = userSlice.actions;
