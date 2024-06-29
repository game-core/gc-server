import {defineStore} from "pinia";

interface AuthState {
    isLogin: boolean;
}

// isLogin ログインしているか
const isLogin = () => {
    const accessToken = useCookie(
        'accessToken',
        {
            secure: true,
            sameSite: true
        }
    );

    return accessToken.value !== null && accessToken.value !== '' && accessToken.value !== undefined;
}

export const useAuthStore = defineStore('auth', {
    state: (): AuthState => ({
        isLogin: isLogin(),
    }),
    actions: {
        setIsLogin(status: boolean): void {
            this.isLogin = status;
        },
    },
});
