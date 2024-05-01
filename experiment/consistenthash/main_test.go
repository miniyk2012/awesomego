package main

import (
	"context"
	"log"
	"testing"

	"github.com/xiaoxuxiansheng/consistent_hash"
	"github.com/xiaoxuxiansheng/consistent_hash/local"
)

func TestLocalConsistentHash(t *testing.T) {
	hashRing := local.NewSkiplistHashRing()
	localMigrator := func(ctx context.Context, dataKeys map[string]struct{}, from, to string) error {
		log.Printf("from: %s, to: %s, data keys: %v", from, to, dataKeys)
		return nil
	}
	//localMigrator = nil
	consistentHash := consistent_hash.NewConsistentHash(hashRing, consistent_hash.NewMurmurHasher(), localMigrator,
		// 每个 node 对应的虚拟节点个数为权重 * replicas
		consistent_hash.WithReplicas(5),
		// 加锁 5 s 后哈希环的锁自动释放
		consistent_hash.WithLockExpireSeconds(5),
	)
	test(t, consistentHash)
}

func test(t *testing.T, consistentHash *consistent_hash.ConsistentHash) {
	ctx := context.Background()
	nodeA := "node_a"
	weightNodeA := 2
	nodeB := "node_b"
	weightNodeB := 1
	nodeC := "node_c"
	weightNodeC := 1
	if err := consistentHash.AddNode(ctx, nodeA, weightNodeA); err != nil {
		t.Error(err)
		return
	}

	if err := consistentHash.AddNode(ctx, nodeB, weightNodeB); err != nil {
		t.Error(err)
		return
	}

	dataKeyA := "data_a"
	dataKeyB := "data_b"
	dataKeyC := "data_c"
	dataKeyD := "data_d"
	node, err := consistentHash.GetNode(ctx, dataKeyA)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyA, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyB); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyB, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyC); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyC, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyD); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyD, node)
	if err := consistentHash.AddNode(ctx, nodeC, weightNodeC); err != nil {
		t.Error(err)
		return
	}
	if node, err = consistentHash.GetNode(ctx, dataKeyA); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyA, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyB); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyB, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyC); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyC, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyD); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyD, node)
	if err = consistentHash.RemoveNode(ctx, nodeB); err != nil {
		t.Error(err)
		return
	}
	if node, err = consistentHash.GetNode(ctx, dataKeyA); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyA, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyB); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyB, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyC); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyC, node)
	if node, err = consistentHash.GetNode(ctx, dataKeyD); err != nil {
		t.Error(err)
		return
	}
	t.Logf("data: %s belongs to node: %s", dataKeyD, node)
	t.Logf("ok")
}
