/**
 * 使用条件变量实现消息队列
 * 编译：g++ -o queue -std=c++11 -lpthread queue.cpp
 * 执行：./queue
 * 执行结果:
 * 62pthread_func=54
    push=1
    push=2
    push=3
    push=4
    push=5
    push wait5
    pop=0x7ffc66cfcf60
    popm=1
    push=6
    pop=0x7ffc66cfcf60
    popm=2
    end
 *
 */
#include <iostream>
#include <queue>
#include <thread>
#include <unistd.h>

typedef int64_t VData;
struct BlockingQueue {
    std::queue<VData> queue_;
    pthread_cond_t full_cond_;
    pthread_cond_t empty_cond_;
    pthread_mutex_t mutex_;
    uint64_t len_;

    int push(VData data) {
        pthread_mutex_lock(&mutex_);
        while (queue_.size() >= len_) { //队列已经满了
            std::cout << "push wait" << queue_.size() << std::endl;
            pthread_cond_wait(&full_cond_, &mutex_);
        }
        std::cout << "push=" << data << std::endl;
        queue_.push(data);
        //signal放在lock内部也可以，放在lock外部也可以。
        //因为pthread_cond_wait是先对mutex_解锁，然后等待信号到来。当信号到来时，会对mutex_进行加锁，加锁成功后pthread_cond_wait退出
        //如果signal在lock内部，则信号到来时，会对pop中mutex_进行加锁，由于mutex_还没有被push退出，因此还需要等待。
        //如果signal在lock外部，则信号到来时，会对pop中mutex_进行加锁，由于mutex_已经被push退出，因此pthread_cond_wait直接退出
        pthread_mutex_unlock(&mutex_);
        pthread_cond_signal(&empty_cond_);
        return 0;
    }
    int pop(VData& data) {
        usleep(1);
        std::cout << "pop=" << &mutex_ << std::endl;
        pthread_mutex_lock(&mutex_);
        while (queue_.size() == 0) {//队列为空
            std::cout << "pop wait"<< std::endl;
            pthread_cond_wait(&empty_cond_, &mutex_);
            std::cout << "pop wait end"<< std::endl;
        }
        data = queue_.front();
        queue_.pop();
        pthread_mutex_unlock(&mutex_);
        pthread_cond_signal(&full_cond_);
        return 0;
    }
};

static void * pthread_func(void *arg) {
    BlockingQueue* bq = (BlockingQueue*)arg;
    VData m;
    std::cout << "pthread_func=54" <<std::endl;
    bq->pop(m);
    std::cout << "popm="<< m << std::endl;
    bq->pop(m);
    std::cout << "popm="<< m << std::endl;

}
int main() {
    BlockingQueue bq;
    bq.len_ = 5;
    bq.mutex_ = PTHREAD_MUTEX_INITIALIZER;
    bq.empty_cond_ = PTHREAD_COND_INITIALIZER;
    bq.full_cond_ = PTHREAD_COND_INITIALIZER;
    VData mm = 1;
    pthread_t tidp;
    if ((pthread_create(&tidp, NULL, pthread_func, (void*)&bq)) == -1) {
        std::cout << "pthread create error" << std::endl;
        return 0;
    }
    VData d;
    std::cout << "62"<< std::endl;
    d = 1;
    bq.push(d);
    d = 2;
    bq.push(d);
    d = 3;
    bq.push(d);
    d = 4;
    bq.push(d);
    d = 5;
    bq.push(d);
    d = 6;
    bq.push(d);
    pthread_join(tidp, NULL);
    std::cout << "end" << std::endl;
}
