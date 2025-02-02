import prisma from 'src/db/db.service';

export async function surveyCreate(
  title: string,
  description: string,
  loginid: string,
): Promise<number> {
  const userid = await prisma.user.findUnique({
    where: {
      loginid: loginid,
    },
    select: {
      id: true,
    },
  });

  const survey = await prisma.survey.create({
    data: {
      title: title,
      description: description,
      userId: userid.id,
    },
  });
  return survey.id;
}
