step1:选举
step2:日志复制，接受leader的共识请求，Follower讲日志发到其他节点进行vote如果成功就执行，并且让其他的节点也同步
Leader Election
时间阀值，谁先跑完，谁就是Candidate，如果两个节点同时成为候选者，则可以进行分组投票。