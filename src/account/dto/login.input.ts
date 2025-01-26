import { InputType, Field } from '@nestjs/graphql';

@InputType()
export class LoginInput {
  @Field()
  loginid: string;

  @Field()
  password: string;
}
