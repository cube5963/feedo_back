import { Auth } from './auth';
import { User } from './user';

const accountService = new User();
const authService = new Auth();

export { accountService, authService };
