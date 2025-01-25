-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,
    "loginid" TEXT NOT NULL,
    "password" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_loginid_key" ON "User"("loginid");
