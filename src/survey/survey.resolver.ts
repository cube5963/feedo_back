import { Mutation, Resolver, Args, Query, Int } from '@nestjs/graphql';
import { SurveyService } from './survey.service';
import { CreateSurveyInput } from './dto/create.dto';
import { Survey } from './dto/get.response';

@Resolver()
export class SurveyResolver {
  constructor(private readonly surveyService: SurveyService) {}

  @Query(() => Survey)
  async getSurvey(@Args('surveyid') surveyid: number) {
    return this.surveyService.Get(surveyid);
  }

  @Mutation(() => Int)
  async createSurvey(
    @Args('createsurveyinput') CreateSurveyInput: CreateSurveyInput,
  ): Promise<number> {
    return await this.surveyService.Create(CreateSurveyInput);
  }
}
