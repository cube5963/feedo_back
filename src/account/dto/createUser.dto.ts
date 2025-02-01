import { Field, InputType } from '@nestjs/graphql';

@InputType()
export class CreateUserInput {
  @Field()
  loginid: string;

  @Field()
  password: string;
}
